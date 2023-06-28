package autoburn

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgAutoburn = "autoburn"
)

type MsgAutoburn struct {
	Amount sdk.Coins      `json:"amount"`
	Sender sdk.AccAddress `json:"sender"`
}

func NewMsgAutoburn(amount sdk.Coins, sender sdk.AccAddress) MsgAutoburn {
	return MsgAutoburn{
		Amount: amount,
		Sender: sender,
	}
}

func (msg MsgAutoburn) Route() string { return ModuleName }

func (msg MsgAutoburn) Type() string { return TypeMsgAutoburn }

func (msg MsgAutoburn) ValidateBasic() error {
	if msg.Amount.Empty() {
		return fmt.Errorf("amount cannot be empty")
	}
	if msg.Sender.Empty() {
		return fmt.Errorf("sender cannot be empty")
	}
	return nil
}

func (msg MsgAutoburn) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgAutoburn) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

type GenesisState struct {
	// define genesis state fields here
}

func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

func (gs GenesisState) Validate() error {
	// validate genesis state fields here
	return nil
}

func InitGenesis(ctx sdk.Context, k Keeper, data GenesisState) {
	// initialize module state from genesis state here
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var data GenesisState
	// export module state to genesis state here
	return data
}
