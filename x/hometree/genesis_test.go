package hometree_test

import (
	"testing"

	keepertest "HomeTree/testutil/keeper"
	"HomeTree/testutil/nullify"
	"HomeTree/x/hometree"
	"HomeTree/x/hometree/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HometreeKeeper(t)
	hometree.InitGenesis(ctx, *k, genesisState)
	got := hometree.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
