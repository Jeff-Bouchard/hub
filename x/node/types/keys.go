package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

const (
	ModuleName     = "node"
	ParamsSubspace = ModuleName
	QuerierRoute   = ModuleName
)

var (
	RouterKey = ModuleName
	StoreKey  = ModuleName
)

var (
	EventModuleName = sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
	)
)

var (
	NodeKeyPrefix = []byte{0x00}

	ActiveNodeKeyPrefix   = []byte{0x10}
	InactiveNodeKeyPrefix = []byte{0x11}

	ActiveNodeForProviderKeyPrefix   = []byte{0x20}
	InactiveNodeForProviderKeyPrefix = []byte{0x21}

	ActiveNodeAtKeyPrefix   = []byte{0x30}
	InactiveNodeAtKeyPrefix = []byte{0x31}
)

func NodeKey(address hub.NodeAddress) []byte {
	return append(NodeKeyPrefix, address.Bytes()...)
}

func ActiveNodeKey(address hub.NodeAddress) []byte {
	return append(ActiveNodeKeyPrefix, address.Bytes()...)
}

func InactiveNodeKey(address hub.NodeAddress) []byte {
	return append(InactiveNodeKeyPrefix, address.Bytes()...)
}

func GetActiveNodeForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append(ActiveNodeForProviderKeyPrefix, address.Bytes()...)
}

func ActiveNodeForProviderKey(p hub.ProvAddress, n hub.NodeAddress) []byte {
	return append(GetActiveNodeForProviderKeyPrefix(p), n.Bytes()...)
}

func GetInactiveNodeForProviderKeyPrefix(address hub.ProvAddress) []byte {
	return append(InactiveNodeForProviderKeyPrefix, address.Bytes()...)
}

func InactiveNodeForProviderKey(p hub.ProvAddress, n hub.NodeAddress) []byte {
	return append(GetInactiveNodeForProviderKeyPrefix(p), n.Bytes()...)
}

func GetActiveNodeAtKeyPrefix(at time.Time) []byte {
	return append(ActiveNodeAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func ActiveNodeAtKey(at time.Time, address hub.NodeAddress) []byte {
	return append(GetActiveNodeAtKeyPrefix(at), address.Bytes()...)
}

func GetInactiveNodeAtKeyPrefix(at time.Time) []byte {
	return append(InactiveNodeAtKeyPrefix, sdk.FormatTimeBytes(at)...)
}

func InactiveNodeAtKey(at time.Time, address hub.NodeAddress) []byte {
	return append(GetInactiveNodeAtKeyPrefix(at), address.Bytes()...)
}
