package internal

import (
	"encoding/json"
	"time"
)

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
