package types

const (
	// ModuleName defines the module name
	ModuleName = "id"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	DefaultParamspace = ModuleName
)

const (
	QueryPathAddress = "address"
	QueryPathId      = "id"
)

// ID message types
const (
	TypeMsgCreateID       = "create_id"
	TypeMsgCreateIDBatch  = "create_id_batch"
	TypeMsgUpdateID       = "update_id"
	TypeMsgReplaceIdOwner = "replace_id_owner"
)

// ID events
const (
	EventCreateID       = "create_id"
	EventCreateIDBatch  = "create_id_batch"
	EventUpdateID       = "update_id"
	EventReplaceIDOwner = "replce_id_owner"
)

const (
	EventAttrIssuer = "issuer"
	EventAttrId     = "id"
	EventAttrOwner  = "owner"
)
