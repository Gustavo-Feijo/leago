package internal

import (
	"encoding/json"
	"time"
)

// UnixMillisTime wraps time.Time and deserialises Riot API timestamps which come in Unix milliseconds (int64).
type UnixMillisTime struct {
	time.Time
}

func (t *UnixMillisTime) UnmarshalJSON(b []byte) error {
	var ms int64
	if err := json.Unmarshal(b, &ms); err != nil {
		return err
	}

	t.Time = time.UnixMilli(ms).UTC()
	return nil
}

// SecondsDuration wraps time.Duration and deserialises Riot API duration when come as plain integer number of seconds.
type SecondsDuration struct {
	time.Duration
}

func (d *SecondsDuration) UnmarshalJSON(b []byte) error {
	var s int64
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	d.Duration = time.Duration(s) * time.Second
	return nil
}

// MillisDuration wraps time.Duration and deserialises Riot API duration when come as plain integer number of milliseconds.
type MillisDuration struct {
	time.Duration
}

func (d *MillisDuration) UnmarshalJSON(b []byte) error {
	var s int64
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	d.Duration = time.Duration(s) * time.Millisecond
	return nil
}

// LobbyTime wraps time.Time and deserialises a custom format used in tournament lobby events.
type LobbyTime struct {
	time.Time
}

func (lt *LobbyTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", s)
	if err != nil {
		return err
	}

	lt.Time = t
	return nil
}
