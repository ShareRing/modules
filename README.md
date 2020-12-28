# Shareledger modules #
Shareledger modules contains all modules for shareledger blockchain.

### Quickstart ###

* Cosmos SDK v0.38.3
* Golang 1.15


### Run test

Run specific testcase
```
go test -v -run TestDuplicatesMsgCreateID
```

Run with coverage result
```
go test -v -cover
```

Run test with detail cover

```
# Run test
go test -coverprofile=coverage.out

# Read coverage result
## Console out put
go tool cover -func=coverage.out
## HTML output
go tool cover -html=coverage.out
```