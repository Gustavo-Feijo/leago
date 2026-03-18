package ratelimit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppKey(t *testing.T) {
	key := AppKey("test")
	assert.Contains(t, key, ":test")
}

func TestMethodKey(t *testing.T) {
	key := MethodKey("test", "method")
	assert.Contains(t, key, "test:method")
}
