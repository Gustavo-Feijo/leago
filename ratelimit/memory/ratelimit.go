package memorylimiter

import (
	"context"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Gustavo-Feijo/leago/ratelimit"
)

type Limit struct {
	count        int
	limit        int
	interval     time.Duration
	mu           sync.Mutex
	lastReset    time.Time
	blockedUntil time.Time
}

func (l *Limit) sync(requests int, interval time.Duration) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.limit == requests && l.interval == interval {
		return false
	}

	l.limit = requests
	l.interval = interval
	return true
}

type Limits []*Limit

func (ls Limits) tryAcquire() time.Duration {
	for _, l := range ls {
		l.mu.Lock()
	}
	defer func() {
		for _, l := range ls {
			l.mu.Unlock()
		}
	}()

	maxWait := ls.computeMaxWait()
	if maxWait > 0 {
		return maxWait
	}

	for _, l := range ls {
		l.count++
	}

	return 0
}

// Go through each limit and get max duration to next refresh.
func (ls Limits) computeMaxWait() time.Duration {
	var maxWait time.Duration
	for _, l := range ls {
		maxWait = maxDuration(maxWait, l.waitDuration())
	}
	return maxWait
}

func (l *Limit) waitDuration() time.Duration {
	if wait := time.Until(l.blockedUntil); wait > 0 {
		return wait
	}
	if time.Since(l.lastReset) >= l.interval {
		l.count = 0
		l.lastReset = time.Now()
	}
	if l.count >= l.limit {
		return l.interval - time.Since(l.lastReset)
	}
	return 0
}

func maxDuration(a, b time.Duration) time.Duration {
	if b > a {
		return b
	}
	return a
}

type MemoryLimiter struct {
	mu             sync.RWMutex
	bucket         map[string][]*Limit
	margin         float64
	intervalMargin time.Duration
}

type Option func(*MemoryLimiter)

// WithLimitSafetyMargin adds a buffer to avoid 429 due to clock skew.
func WithLimitSafetyMargin(margin float64) Option {
	return func(ml *MemoryLimiter) {
		if margin > 1.0 {
			ml.margin = 1.0
			return
		}

		if margin < 0 {
			margin = 0
		}

		ml.margin = margin
	}
}

// WithIntervalSafetyMargin adds a buffer to the reset duration (If resets every 10s, starts to reset every 10s + Xs).
func WithIntervalSafetyMargin(duration time.Duration) Option {
	return func(ml *MemoryLimiter) {
		ml.intervalMargin = duration
	}
}

func NewMemoryLimiter(opts ...Option) ratelimit.RateLimiter {
	limiter := &MemoryLimiter{
		bucket:         map[string][]*Limit{},
		margin:         1.0,
		intervalMargin: 0,
	}

	for _, opt := range opts {
		opt(limiter)
	}

	return limiter
}

func (m *MemoryLimiter) Acquire(ctx context.Context, appKey, methodKey string) error {
	appLimits := m.getOrCreate(appKey)
	methodLimits := m.getOrCreate(methodKey)

	// Suppress linting warning that requires append to be reassigned to the slice.
	//nolint:gocritic
	combined := append(appLimits, methodLimits...)

	for {
		wait := combined.tryAcquire()
		if wait == 0 {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(wait):
		}
	}
}

func (m *MemoryLimiter) getOrCreate(key string) Limits {
	m.mu.Lock()
	defer m.mu.Unlock()
	if b, ok := m.bucket[key]; ok {
		return b
	}

	b := Limits{
		// Using default limits for development keys.
		{limit: int(math.Floor(20 * m.margin)), interval: (time.Second) + m.intervalMargin, lastReset: time.Now()},
		{limit: int(math.Floor(100 * m.margin)), interval: (2 * time.Minute) + m.intervalMargin, lastReset: time.Now()},
	}

	m.bucket[key] = b
	return b
}

func (m *MemoryLimiter) Sync(_ context.Context, headers http.Header, appKey, methodKey string) {
	if appHeader := headers.Get("X-App-Rate-Limit"); appHeader != "" {
		m.syncKey(appKey, appHeader)
	}
	if methodHeader := headers.Get("X-Method-Rate-Limit"); methodHeader != "" {
		m.syncKey(methodKey, methodHeader)
	}
}

func (m *MemoryLimiter) syncKey(key, raw string) {
	entries := strings.Split(raw, ",")

	m.mu.Lock()
	defer m.mu.Unlock()
	margin := m.margin
	intervalMargin := m.intervalMargin
	limits, ok := m.bucket[key]
	if !ok {
		return
	}
	for len(limits) < len(entries) {
		limits = append(limits, &Limit{lastReset: time.Now()})
	}
	m.bucket[key] = limits

	for i, entry := range entries {
		parts := strings.SplitN(entry, ":", 2)
		if len(parts) != 2 {
			continue
		}
		requests, err1 := strconv.Atoi(parts[0])
		secs, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil || secs == 0 {
			continue
		}

		// At least 1 req allowed.
		requestsWithMargin := math.Max(1, math.Floor(margin*float64(requests)))

		limits[i].sync(int(requestsWithMargin), (time.Duration(secs)*time.Second)+intervalMargin)
	}
}

func (m *MemoryLimiter) NotifyTooManyRequests(ctx context.Context, headers http.Header, appKey, methodKey string) {
	retryAfter := parseRetryAfter(headers)
	if retryAfter == 0 {
		return
	}

	blockedUntil := time.Now().Add(retryAfter)

	for _, key := range []string{appKey, methodKey} {
		m.mu.RLock()
		limits, ok := m.bucket[key]
		m.mu.RUnlock()
		if !ok {
			continue
		}
		for _, l := range limits {
			l.mu.Lock()
			if blockedUntil.After(l.blockedUntil) {
				l.blockedUntil = blockedUntil
			}
			l.mu.Unlock()
		}
	}
}

func parseRetryAfter(headers http.Header) time.Duration {
	val := headers.Get("Retry-After")
	if val == "" {
		return 0
	}
	secs, err := strconv.Atoi(val)
	if err != nil || secs <= 0 {
		return 0
	}
	return time.Duration(secs) * time.Second
}
