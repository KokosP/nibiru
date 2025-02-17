package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	// FlagPoolFile Will be parsed to string.
	FlagPoolFile = "pool-file"

	// FlagPoolId Will be parsed to uint64.
	FlagPoolId = "pool-id"

	// FlagUseAllCoins Will be parsed to uint64.
	FlagUseAllCoins = "use-all-coins"

	// FlagTokensIn Will be parsed to []sdk.Coin.
	FlagTokensIn = "tokens-in"

	// FlagPoolSharesOut Will be parsed to sdk.Coin.
	FlagPoolSharesOut = "pool-shares-out"

	// FlagTokenIn Will be parsed to sdk.Coin.
	FlagTokenIn = "token-in"

	// FlagTokenOutDenom Will be parsed to string.
	FlagTokenOutDenom = "token-out-denom"
)

type createPoolInputs struct {
	Weights        string `json:"weights"`
	InitialDeposit string `json:"initial-deposit"`
	SwapFee        string `json:"swap-fee"`
	ExitFee        string `json:"exit-fee"`
}

func FlagSetCreatePool() *flag.FlagSet {
	fs := flag.NewFlagSet("create-pool", flag.PanicOnError)

	fs.String(FlagPoolFile, "", "Pool json file path")
	return fs
}

func FlagSetJoinPool() *flag.FlagSet {
	fs := flag.NewFlagSet("join-pool", flag.PanicOnError)

	fs.Uint64(FlagPoolId, 0, "The id of pool")
	fs.StringArray(FlagTokensIn, []string{""}, "Amount of each denom to send into the pool (specify multiple denoms with: --tokens-in=1uusdc --tokens-in=1unusd)")
	fs.Bool(FlagUseAllCoins, false, "Whether to use all the tokens in tokens-in to maximize shares out with a swap first")
	return fs
}

func FlagSetExitPool() *flag.FlagSet {
	fs := flag.NewFlagSet("exit-pool", flag.ContinueOnError)

	fs.Uint64(FlagPoolId, 0, "The pool id to withdraw from.")
	fs.String(FlagPoolSharesOut, "", "The amount of pool share tokens to burn.")
	return fs
}

func FlagSetSwapAssets() *flag.FlagSet {
	fs := flag.NewFlagSet("swap-assets", flag.ContinueOnError)

	fs.Uint64(FlagPoolId, 0, "The pool id to withdraw from.")
	fs.String(FlagTokenIn, "", "The amount of tokens to swap in.")
	fs.String(FlagTokenOutDenom, "", "The denom of the token to extract.")
	return fs
}
