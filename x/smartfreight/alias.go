package smartfreight

import (
	"smartfreight/x/smartfreight/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgSetDelivery      = types.NewMsgSetDelivery
	NewMsgCompleteDelivery = types.NewMsgCompleteDelivery
	NewDelivery            = types.NewDelivery
	ModuleCdc              = types.ModuleCdc
	RegisterCodec          = types.RegisterCodec
)

type (
	MsgSetDelivery      = types.MsgSetDelivery
	MsgCompleteDelivery = types.MsgCompleteDelivery
	QueryResResolve     = types.QueryResResolve
	QueryResKEys        = types.QueryResKeys
	Delivery            = types.Delivery
)
