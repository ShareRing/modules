# Id modules

## State
## Messages
## Query
### Create new id
Create new id by given information
```
slcli tx id create <id> <backup_address> <owner_address> <extra_data>
```
### Get id info
It supports querying by address or by id.
```
slcli query id info address <address> 
slcli query id info id <id>
```