package autoburn

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) HandleMsg(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
	switch msg := msg.(type) {
	case *MsgAutoburn:
		return k.handleAutoburn(ctx, msg)
	default:
		return nil, errorsmod.Wrapf(errortypes.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
	}
}

func (k Keeper) handleAutoburn(ctx sdk.Context, msg *MsgAutoburn) (*sdk.Result, error) {
	err := k.BurnCoins(ctx, msg.Amount, msg.Sender)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{
		Log: fmt.Sprintf("Successfully burned %s from %s", msg.Amount.String(), msg.Sender.String()),
	}, nil
}
