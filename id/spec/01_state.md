<!--
order: 1
-->

# State

## Id
- Address: `<0x1> | <address> -> amino(id)`

- ID detail: `<0x2> | <id> -> amino(DocDetailState)`

    ```go
    type BaseID struct {
        IssuerAddr sdk.AccAddress `json:"issuer_address"`
        BackupAddr sdk.AccAddress `json:"backup_address"`
        OwnerAddr  sdk.AccAddress `json:"owner_address"`
        ExtraData  string         `json:"extra_data"`
    }
    ```

