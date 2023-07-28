## Fabric Shell Script - network

## Hyperledger Fabric v2.2

It is simply for building a blockchain environment, and `chaincode` and `client` need to be built separately.

- 1 orderer
- 1 peer & couchdb
- (optional) 1 fabric-ca

Start the Fabric network with the following command

```sh
# with fabric-ca & client for API test
./network.sh start

# clear docker
./network.sh clear

# restart with fabric-ca
./network.sh restart

```
