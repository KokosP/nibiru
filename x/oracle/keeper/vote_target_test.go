package keeper

import (
	"testing"

	"github.com/NibiruChain/nibiru/collections"

	"github.com/stretchr/testify/require"
)

func TestKeeper_GetVoteTargets(t *testing.T) {
	input := CreateTestInput(t)

	for _, p := range input.OracleKeeper.Pairs.Iterate(input.Ctx, collections.Range[string]{}).Keys() {
		input.OracleKeeper.Pairs.Delete(input.Ctx, p)
	}

	expectedTargets := []string{"bar", "foo", "whoowhoo"}
	for _, target := range expectedTargets {
		input.OracleKeeper.Pairs.Insert(input.Ctx, target)
	}

	targets := input.OracleKeeper.GetVoteTargets(input.Ctx)
	require.Equal(t, expectedTargets, targets)
}

func TestKeeper_IsVoteTarget(t *testing.T) {
	input := CreateTestInput(t)

	for _, p := range input.OracleKeeper.Pairs.Iterate(input.Ctx, collections.Range[string]{}).Keys() {
		input.OracleKeeper.Pairs.Delete(input.Ctx, p)
	}

	validTargets := []string{"bar", "foo", "whoowhoo"}
	for _, target := range validTargets {
		input.OracleKeeper.Pairs.Insert(input.Ctx, target)
		require.True(t, input.OracleKeeper.IsVoteTarget(input.Ctx, target))
	}
}
