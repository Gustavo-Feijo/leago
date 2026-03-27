package profileicon

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntOrString(t *testing.T) {
	tests := []struct {
		name        string
		input       []byte
		expected    int
		expectError bool
	}{
		{
			name:     "valid string number",
			input:    []byte(`"15"`),
			expected: 15,
		},
		{
			name:     "valid number",
			input:    []byte(`15`),
			expected: 15,
		},
		{
			name:        "invalid numeric string",
			input:       []byte(`"abc"`),
			expectError: true,
		},
		{
			name:        "valid json invalid type",
			input:       []byte(`true`),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result IntOrString

			err := json.Unmarshal([]byte(tt.input), &result)

			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

			assert.Equal(t, IntOrString(tt.expected), result)
		})
	}
}
