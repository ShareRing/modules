package types

const (
	// ModuleName defines the module name
	ModuleName = "doc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	DefaultParamspace = ModuleName
)
const (
	TypeMsgCreateDoc        = "create_doc"
	TypeMsgCreateDocInBatch = "create_doc_batch"
	TypeMsgUpdateDoc        = "update_doc"
	TypeMsgRevokeDoc        = "revoke_doc"
)

const (
	QueryByProof = "proof"
)
const (
	EventTypeCreateDoc = "create_doc"

	EventAttrHolder = "hodler"
	EventAttrProof  = "proof"
	EventAttrIssuer = "issuer"
)

var (
	StateKeySep     = []byte("|")
	DocDetailPrefix = []byte{0x1}
	DocBasicPrefix  = []byte{0x2}
)
