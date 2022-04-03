package keeper_test

import (
	"testing"

	"github.com/MatrixDao/matrix/x/stablecoin/keeper"
	"github.com/MatrixDao/matrix/x/testutil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestAsInt(t *testing.T) {
	testCases := []struct {
		name   string
		inDec  sdk.Dec
		outInt sdk.Int
	}{
		{
			name:   "One to int",
			inDec:  sdk.NewDec(1),
			outInt: sdk.NewInt(1),
		},
		{
			name:   "Small loss of precision due to truncation",
			inDec:  sdk.MustNewDecFromStr("1.1"),
			outInt: sdk.NewInt(1),
		},
		{
			name:   "Large loss of precision due to truncation",
			inDec:  sdk.MustNewDecFromStr("9.999"),
			outInt: sdk.NewInt(9),
		},
	}
	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			sdkInt := keeper.AsInt(tc.inDec)
			require.Equal(t, tc.outInt, sdkInt)
		})
	}

}

func TestMint_NeededCollAmtGivenGov(t *testing.T) {
	testCases := []struct {
		name              string
		govAmt            sdk.Int
		priceGov          sdk.Dec
		priceColl         sdk.Dec
		collRatio         sdk.Dec
		neededCollAmt     sdk.Int
		mintableStableAmt sdk.Int
		err               error
	}{
		{
			name:              "Low collateral ratio",
			govAmt:            sdk.NewInt(10),
			priceGov:          sdk.NewDec(90),
			priceColl:         sdk.NewDec(10),
			collRatio:         sdk.MustNewDecFromStr("0.1"),
			neededCollAmt:     sdk.NewInt(9),
			mintableStableAmt: sdk.NewInt(999),
			err:               nil,
		}, {
			name:              "High collateral ratio",
			govAmt:            sdk.NewInt(10),
			priceGov:          sdk.MustNewDecFromStr("1"),
			priceColl:         sdk.MustNewDecFromStr("2"),
			collRatio:         sdk.MustNewDecFromStr("0.9"),
			neededCollAmt:     sdk.NewInt(45),
			mintableStableAmt: sdk.NewInt(100),
			err:               nil,
		},
	}
	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			neededCollAmt, mintableStableAmt := keeper.NeededCollAmtGivenGov(
				tc.govAmt, tc.priceGov, tc.priceColl, tc.collRatio)
			testutil.RequireEqualWithMessage(
				t, neededCollAmt, tc.neededCollAmt, "neededCollAmt")
			testutil.RequireEqualWithMessage(
				t, mintableStableAmt, tc.mintableStableAmt, "mintableStableAmt")
		})
	}
}

func TestMint_NeededGovAmtGivenColl(t *testing.T) {
	testCases := []struct {
		name              string
		collAmt           sdk.Int
		priceGov          sdk.Dec
		priceColl         sdk.Dec
		collRatio         sdk.Dec
		neededGovAmt      sdk.Int
		mintableStableAmt sdk.Int
		err               error
	}{
		{
			name:              "collRatio above 50%",
			collAmt:           sdk.NewInt(70),
			priceGov:          sdk.NewDec(10),
			priceColl:         sdk.NewDec(1), // 70 * 1 = 70
			collRatio:         sdk.MustNewDecFromStr("0.7"),
			neededGovAmt:      sdk.NewInt(3), // 10 * 3 = 30
			mintableStableAmt: sdk.NewInt(100),
			err:               nil,
		}, {
			name:              "collRatio below 50%",
			collAmt:           sdk.NewInt(40),
			priceGov:          sdk.NewDec(10),
			priceColl:         sdk.NewDec(2),
			collRatio:         sdk.MustNewDecFromStr("0.8"),
			neededGovAmt:      sdk.NewInt(2),
			mintableStableAmt: sdk.NewInt(100),
			err:               nil,
		},
	}
	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			neededGovAmt, mintableStableAmt := keeper.NeededGovAmtGivenColl(
				tc.collAmt, tc.priceGov, tc.priceColl, tc.collRatio)
			require.Equal(t, neededGovAmt, tc.neededGovAmt)
			require.Equal(t, mintableStableAmt, tc.mintableStableAmt)
		})
	}
}