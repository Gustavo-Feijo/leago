package memorylimiter

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMemoryLimiter(t *testing.T) {
	rl := NewMemoryLimiter(
		WithIntervalSafetyMargin(time.Second),
		WithLimitSafetyMargin(0.7),
	).(*MemoryLimiter)

	assert.NotNil(t, rl)
	assert.Equal(t, time.Second, rl.intervalMargin)
	assert.Equal(t, 0.7, rl.margin)
}

func TestNewMemoryLimiterNegativeMargin(t *testing.T) {
	rl := NewMemoryLimiter(
		WithLimitSafetyMargin(-1.0),
	).(*MemoryLimiter)

	assert.NotNil(t, rl)
	assert.Equal(t, 0.0, rl.margin)
}

func TestNewMemoryLimiterMarginTooBig(t *testing.T) {
	rl := NewMemoryLimiter(
		WithLimitSafetyMargin(1.2),
	).(*MemoryLimiter)

	assert.NotNil(t, rl)
	assert.Equal(t, 1.0, rl.margin)
}

func TestLimitSync(t *testing.T) {
	l := &Limit{}

	changed := l.sync(10, time.Second)
	assert.True(t, changed)

	changed = l.sync(10, time.Second)
	assert.False(t, changed)
}

func TestLimitsTryAcquire(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name string

		limits Limits

		expected bool
	}{
		{
			name: "allowed immediately",

			limits: Limits{
				{limit: 10, count: 0, interval: time.Second, lastReset: now},
			},

			expected: false,
		},
		{
			name: "blocked by count",

			limits: Limits{
				{limit: 1, count: 1, interval: time.Second, lastReset: now},
			},

			expected: true,
		},
		{
			name: "blocked by blockedUntil",

			limits: Limits{
				{limit: 10, count: 0, interval: time.Second, lastReset: now, blockedUntil: now.Add(time.Second)},
			},

			expected: true,
		},
		{
			name: "reset after interval",

			limits: Limits{
				{limit: 1, count: 1, interval: time.Millisecond, lastReset: now.Add(-time.Second)},
			},

			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wait := tt.limits.tryAcquire()

			if tt.expected {
				assert.True(t, wait > 0)
			} else {
				assert.Equal(t, time.Duration(0), wait)
			}
		})
	}
}

func TestAcquire(t *testing.T) {
	ml := NewMemoryLimiter()

	ctx := context.Background()

	err := ml.Acquire(ctx, "app", "method")
	assert.NoError(t, err)
}

func TestAcquireWait(t *testing.T) {
	ml := NewMemoryLimiter().(*MemoryLimiter)

	limits := ml.getOrCreate("app")
	limits[0].interval = time.Second
	limits[0].limit = 1
	limits[0].count = 1

	ctx := context.Background()
	err := ml.Acquire(ctx, "app", "method")
	assert.NoError(t, err)

	err = ml.Acquire(ctx, "app", "method")
	assert.NoError(t, err)
}

func TestAcquireContextCancel(t *testing.T) {
	ml := NewMemoryLimiter().(*MemoryLimiter)

	limits := ml.getOrCreate("app")
	limits[0].limit = 1
	limits[0].count = 1

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := ml.Acquire(ctx, "app", "method")
	assert.Error(t, err)
}

func TestGetOrCreate(t *testing.T) {
	ml := NewMemoryLimiter().(*MemoryLimiter)

	l1 := ml.getOrCreate("key")
	l2 := ml.getOrCreate("key")

	assert.Equal(t, l1, l2)
}

func TestSync(t *testing.T) {
	ml := NewMemoryLimiter().(*MemoryLimiter)

	headers := http.Header{}
	headers.Set("X-App-Rate-Limit", "10:1,20:2")
	headers.Set("X-Method-Rate-Limit", "5:1")

	ml.getOrCreate("app")
	ml.getOrCreate("method")

	ml.Sync(context.Background(), headers, "app", "method")

	appLimits := ml.bucket["app"]
	assert.Equal(t, 2, len(appLimits))
	assert.Equal(t, 10, appLimits[0].limit)
	assert.Equal(t, 20, appLimits[1].limit)
}

func TestSyncExpandLimits(t *testing.T) {
	ml := NewMemoryLimiter().(*MemoryLimiter)

	headers := http.Header{}
	headers.Set("X-App-Rate-Limit", "10:1,20:2,30:3,40:4")

	ml.getOrCreate("app")

	ml.Sync(context.Background(), headers, "app", "method")

	assert.Equal(t, 4, len(ml.bucket["app"]))
}

func TestSyncKeyKeyDoesNotExist(t *testing.T) {
	ml := NewMemoryLimiter().(*MemoryLimiter)

	ml.syncKey("missing", "10:1")
}

func TestSyncInvalidHeader(t *testing.T) {
	ml := NewMemoryLimiter().(*MemoryLimiter)

	headers := http.Header{}
	headers.Set("X-App-Rate-Limit", "invalid")

	ml.getOrCreate("app")

	ml.Sync(context.Background(), headers, "app", "method")
}

func TestSyncInvalidEntries(t *testing.T) {
	ml := NewMemoryLimiter().(*MemoryLimiter)

	headers := http.Header{}
	headers.Set("X-App-Rate-Limit", "invalid,10:abc,20:0")

	ml.getOrCreate("app")

	ml.Sync(context.Background(), headers, "app", "method")

	assert.True(t, true)
}

func TestNotifyTooManyRequests(t *testing.T) {
	ml := NewMemoryLimiter().(*MemoryLimiter)

	headers := http.Header{}
	headers.Set("Retry-After", "1")

	ml.getOrCreate("app")
	ml.getOrCreate("method")

	ml.NotifyTooManyRequests(context.Background(), headers, "app", "method")

	for _, l := range ml.bucket["app"] {
		assert.True(t, l.blockedUntil.After(time.Now()))
	}
}

func TestNotifyTooManyRequestsNoHeader(t *testing.T) {
	ml := NewMemoryLimiter().(*MemoryLimiter)

	headers := http.Header{}

	ml.NotifyTooManyRequests(context.Background(), headers, "app", "method")

	assert.True(t, true)
}

func TestNotifyTooManyRequestsMissingKey(t *testing.T) {
	ml := NewMemoryLimiter().(*MemoryLimiter)

	headers := http.Header{}
	headers.Set("Retry-After", "1")

	ml.NotifyTooManyRequests(context.Background(), headers, "app", "method")
}

func TestParseRetryAfter(t *testing.T) {
	tests := []struct {
		header   string
		expected time.Duration
	}{
		{"10", 10 * time.Second},
		{"0", 0},
		{"invalid", 0},
		{"", 0},
	}

	for _, tt := range tests {
		h := http.Header{}
		h.Set("Retry-After", tt.header)

		result := parseRetryAfter(h)
		assert.Equal(t, tt.expected, result)
	}
}
