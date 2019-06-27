package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

type MsgCompleteDelivery struct {
	JobID  int64          `json:"jobId"`
	Broker sdk.AccAddress `json:"broker"`
}

type MsgSetDelivery struct {
	JobID    int64    `json:"jobId"`
	Delivery Delivery `json:"delivery"`
}

func NewMsgSetDelivery(JobID int64, Delivery Delivery) MsgSetDelivery {
	return MsgSetDelivery{
		JobID,
		Delivery,
	}
}
func NewMsgCompleteDelivery(JobID int64, Broker sdk.AccAddress) MsgCompleteDelivery {
	return MsgCompleteDelivery{
		JobID,
		Broker,
	}
}
func (msg MsgSetDelivery) Route() string      { return RouterKey }
func (msg MsgCompleteDelivery) Route() string { return RouterKey }

func (msg MsgSetDelivery) Type() string      { return "set_delivery" }
func (msg MsgCompleteDelivery) Type() string { return "complete_delivery" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetDelivery) ValidateBasic() sdk.Error {
	if msg.Delivery.Shipper.Empty() {
		return sdk.ErrInvalidAddress(msg.Delivery.Shipper.String())
	}
	if msg.Delivery.Broker.Empty() {
		return sdk.ErrInvalidAddress(msg.Delivery.Broker.String())
	}
	if msg.Delivery.Price.Empty() {
		return sdk.ErrUnknownRequest("Price cannot be empty")
	}
	if msg.Delivery.Completed == true {
		return sdk.ErrUnknownRequest("Cannot set completed delivery")
	}
	return nil
}

func (msg MsgCompleteDelivery) ValidateBasic() sdk.Error {
	if msg.JobID > 0 {
		return nil
	}
	return sdk.ErrUnknownRequest("JobID is not valid")
}

func (msg MsgSetDelivery) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgCompleteDelivery) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetDelivery) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Delivery.Shipper}
}

func (msg MsgCompleteDelivery) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Broker}
}
