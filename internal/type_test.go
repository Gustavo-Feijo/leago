package internal

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnixMillisTimeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    time.Time
		expectError bool
	}{
		{
			name:     "valid timestamp",
			input:    `1700000000000`,
			expected: time.UnixMilli(1700000000000).UTC(),
		},
		{
			name:        "invalid timestamp",
			input:       `"not-a-number"`,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result UnixMillisTime

			err := json.Unmarshal([]byte(tt.input), &result)

			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			assert.True(t, result.Equal(tt.expected))
		})
	}
}

func TestSecondsDurationUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    time.Duration
		expectError bool
	}{
		{
			name:     "valid duration",
			input:    `1800`,
			expected: time.Duration(time.Minute * 30),
		},
		{
			name:        "invalid duration",
			input:       `"not-a-number"`,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result SecondsDuration

			err := json.Unmarshal([]byte(tt.input), &result)

			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			assert.Equal(t, tt.expected.Abs(), result.Abs())
		})
	}
}
