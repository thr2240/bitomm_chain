package autoburn

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/gogo/protobuf/codec"
)

type Keeper struct {
	storeKey   storetypes.StoreKey
	cdc        *codec.Codec
	bankKeeper bankkeeper.Keeper
}

func NewKeeper(storeKey storetypes.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

func (k Keeper) BurnCoins(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error {
	err := k.bankKeeper.SendCoins(ctx, sender, sdk.AccAddress{}, amount)
	if err != nil {
		return err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			TypeMsgAutoburn,
			sdk.NewAttribute("sender", sender.String()),
			sdk.NewAttribute("amount", amount.String()),
		),
	)
	return nil
}
