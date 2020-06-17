package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgRegisterProvider)(nil)
	_ sdk.Msg = (*MsgUpdateProvider)(nil)
)

type MsgRegisterProvider struct {
	From        sdk.AccAddress `json:"from"`
	Name        string         `json:"name"`
	Website     string         `json:"website"`
	Description string         `json:"description"`
}

func NewMsgRegisterProvider(from sdk.AccAddress, name, website, description string) MsgRegisterProvider {
	return MsgRegisterProvider{
		From:        from,
		Name:        name,
		Website:     website,
		Description: description,
	}
}

func (m MsgRegisterProvider) Route() string {
	return RouterKey
}

func (m MsgRegisterProvider) Type() string {
	return "register_provider"
}

func (m MsgRegisterProvider) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}
	if len(m.Name) == 0 || len(m.Name) > 32 {
		return ErrorInvalidField("name")
	}
	if len(m.Website) > 32 {
		return ErrorInvalidField("website")
	}
	if len(m.Description) > 256 {
		return ErrorInvalidField("description")
	}

	return nil
}

func (m MsgRegisterProvider) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRegisterProvider) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}

type MsgUpdateProvider struct {
	From        sdk.AccAddress `json:"from"`
	ID          hub.ProviderID `json:"id"`
	Name        string         `json:"name"`
	Website     string         `json:"website"`
	Description string         `json:"description"`
}

func (m MsgUpdateProvider) Route() string {
	return RouterKey
}

func (m MsgUpdateProvider) Type() string {
	return "update_provider"
}

func (m MsgUpdateProvider) ValidateBasic() sdk.Error {
	if m.From == nil || m.From.Empty() {
		return ErrorInvalidField("from")
	}
	if m.ID == nil {
		return ErrorInvalidField("id")
	}
	if len(m.Name) != 0 && len(m.Name) > 32 {
		return ErrorInvalidField("name")
	}
	if len(m.Website) != 0 && len(m.Website) > 32 {
		return ErrorInvalidField("website")
	}
	if len(m.Description) != 0 && len(m.Description) > 256 {
		return ErrorInvalidField("description")
	}

	return nil
}

func (m MsgUpdateProvider) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpdateProvider) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}
