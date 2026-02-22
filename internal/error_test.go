package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRiotError(t *testing.T) {

	riotError := RiotError{
		StatusCode: 400,
		Status:     "error",
		Body:       "payload",
	}

	assert.Contains(t, riotError.Error(), "body:")

	riotErrorNoBody := RiotError{
		StatusCode: 500,
		Status:     "error",
	}
	assert.NotContains(t, riotErrorNoBody.Error(), "body:")
}
