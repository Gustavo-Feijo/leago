package options_test

import (
	"leago/internal"
	"leago/options"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeOptions_Length(t *testing.T) {
	merged := options.MergeOptions(
		[]internal.RequestOption{
			internal.WithApiMethod("Default"),
		},
		[]options.PublicOption{
			options.WithApiMethod("Override"),
		},
	)

	require.Len(t, merged, 2)
}
