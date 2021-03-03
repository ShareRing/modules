<!--
order: 3
-->
# Messages

## MsgCreateId
```go
type MsgCreateId struct {
	BackupAddr sdk.AccAddress `json:"backup_address"`
	ExtraData  string         `json:"extra_data"`
	Id         string         `json:"id"`
	IssuerAddr sdk.AccAddress `json:"issuer_address"`
	OwnerAddr  sdk.AccAddress `json:"owner_address"`
}
```
## MsgCreateIdBatch
```go
type MsgCreateIdBatch struct {
	BackupAddr []sdk.AccAddress `json:"backup_address"`
	ExtraData  []string         `json:"extra_data"`
	Id         []string         `json:"id"`
	IssuerAddr sdk.AccAddress   `json:"issuer_address"`
	OwnerAddr  []sdk.AccAddress `json:"owner_address"`
}
```

## MsgReplaceIdOwner
```go
type MsgReplaceIdOwner struct {
	BackupAddr sdk.AccAddress `json:"backup_address"`
	Id         string         `json:"id"`
	OwnerAddr  sdk.AccAddress `json:"owner_address"`
}
```

## MsgUpdateId
```go
type MsgUpdateId struct {
	ExtraData  string         `json:"extra_data"`
	Id         string         `json:"id"`
	IssuerAddr sdk.AccAddress `json:"issuer_address"`
}
```