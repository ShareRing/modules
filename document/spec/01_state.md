<!--
order: 1
-->

# State

## Document
- Document detail: `<0x1> | <holder> | <document proof> | <issuer> -> amino(DocDetailState)`

    * `<version>` is a 2 bytes flag.
        * It automatically increases by 1 when the document is updated.
        * The start value is 1.
        * Set the version to 0xFFFF when it’s revoked.


    ```go
    type DocDetailState struct {
        Data    string
        Version uint16
    }
    ```
- Document basic: `<0x2> | <document proof> -> amino(DocBasicState)
`
    ```go
    type DocBasicState struct {
        Holder string
        Issuer sdk.AccAddress
    }
    ```