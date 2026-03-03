package options_test

import (
	"testing"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"

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
