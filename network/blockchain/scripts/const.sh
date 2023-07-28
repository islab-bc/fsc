# src directory path
sdir=$shdir/src/

# build directory path 😊
bdir=$shdir/build

# fabric-version
FABRIC_VERSION=2.2.2

# for orderer
ORDERER_ADDR=orderer0.example.com:7050
GLOBAL_FLAGS="-o $ORDERER_ADDR --tls --cafile /etc/hyperledger/fabric/orderer-tls/tlsca.example.com-cert.pem"
ORGANIZATION=btp

#for peer
BTP_PEER=peer0.btp.example.com

#for channel
CHANNEL=tester-channel

#for chaincode, fill chaincodes with chaincode names
CHAINCODE_DIR=/etc/hyperledger/fabric/chaincodes
CHAINCODES=(board user)