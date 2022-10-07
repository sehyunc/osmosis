package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// constants
const (
	TypeMsgCreateDenom = "create_denom"
	TypeMsgMint        = "mint"
)

var _ sdk.Msg = &MsgCreateDenom{}

// NewMsgCreateValidatorSetPreference creates a msg to create a validator-set preference.
func NewMsgCreateDenom(sender sdk.AccAddress, denomId, denomName, denomData string) *MsgCreateDenom {
	return &MsgCreateDenom{
		Id:        denomId,
		Sender:    sender.String(),
		DenomName: denomName,
		Data:      denomData,
	}
}

func (m MsgCreateDenom) Route() string { return RouterKey }
func (m MsgCreateDenom) Type() string  { return TypeMsgCreateDenom }
func (m MsgCreateDenom) ValidateBasic() error {
	return nil
}

func (m MsgCreateDenom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners takes a create validator-set message and returns the delegator in a byte array.
func (m MsgCreateDenom) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgMint{}

func NewMsgMint(sender sdk.AccAddress, denomId string, amount sdk.Coin) *MsgMint {
	return &MsgMint{
		Id:     denomId,
		Sender: sender.String(),
		Amount: amount,
	}
}

func (m MsgMint) Route() string { return RouterKey }
func (m MsgMint) Type() string  { return TypeMsgMint }
func (m MsgMint) ValidateBasic() error {
	return nil
}

func (m MsgMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

// GetSigners takes a create validator-set message and returns the delegator in a byte array.
func (m MsgMint) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}
