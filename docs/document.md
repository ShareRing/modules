
## New doc
```
./build/slcli tx doc create <holder id> <doc proof> <data>

```

```
 ./build/slcli tx doc create id13 3333 data --from issuer --fees 1shr --yes --broadcast-mode block
```

## New doc batch
```
./slcli tx doc create-batch <holder id 1,holder id2> <doc proof 1,doc proof 2> <data1,data2> --from issuer --fees=1shr --yes --broadcast-mode=block
```

```
./slcli tx doc create-batch "id1","id2" "p1","p2" d1,d2 --from issuer --fees=1shr --yes --broadcast-mode=block
```

## Get doc
```
slcli query doc proof <proof>
slcli query doc holder <holder>
```