package keeper_test

import (
	"testing"

	testkeeper "HomeTree/testutil/keeper"
	"HomeTree/x/hometree/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.HometreeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
