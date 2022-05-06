package perp

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/NibiruChain/nibiru/x/perp/keeper"
	types "github.com/NibiruChain/nibiru/x/perp/types/v1"
)

/*
NewHandler returns an sdk.Handler for "x/perp" messages.
A handler defines the core state transition functions of an application.
First, the handler performs stateful checks to make sure each 'msg' is valid.
At this stage, the 'msg.ValidateBasic()' method has already been called, meaning
stateless checks on the message (like making sure parameters are correctly
formatted) have already been performed.
*/
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgFoo:
			res, err := msgServer.Foo(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			errMsg := fmt.Sprintf(
				"unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}