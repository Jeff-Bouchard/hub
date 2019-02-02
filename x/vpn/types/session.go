package types

import (
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type SessionDetails struct {
	ID            string
	NodeID        string
	ClientAddress csdkTypes.AccAddress
	LockedAmount  csdkTypes.Coin
	PricePerGB    csdkTypes.Coin
	Bandwidth     struct {
		ToProvide     sdkTypes.Bandwidth
		Consumed      sdkTypes.Bandwidth
		NodeOwnerSign []byte
		ClientSign    []byte
		UpdatedAt     time.Time
	}
	Status    string
	StatusAt  time.Time
	StartedAt time.Time
	EndedAt   time.Time
}
