package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// query endpoints supported by the Id Querier
const (
	QueryInfo = "info"
)

// QueryIdByAddressParams defines the params for querying an account id information.
type QueryIdByAddressParams struct {
	Address sdk.AccAddress
}

func NewQueryIdByAddressParams(ownerAddr sdk.AccAddress) QueryIdByAddressParams {
	return QueryIdByAddressParams{Address: ownerAddr}
}

type QueryIdByIdParams struct {
	Id string
}

func NewQueryIdByIdParams(id string) QueryIdByIdParams {
	return QueryIdByIdParams{Id: id}
}
