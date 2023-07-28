#!bin/sh

shdir="$( cd "$(dirname "$0")" ; pwd -P)"

source $shdir/scripts/utils.sh
export COMPOSE_IGNORE_ORPHANS=True

function blockchain_build_cryptogen {
    command "docker run -it --rm \
    -v $bdir:/workdir \
    --workdir /workdir \
    hyperledger/fabric-tools:$FABRIC_VERSION \
    cryptogen generate --config=./asset/tool/crypto-config.yaml --output=./asset/artifacts/crypto-config"
}

function blockchain_build_configtxgen {
    blockchain_build_configtxgen_genesis_block
    blockchain_build_configtxgen_channel_tx $CHANNEL
}

function blockchain_build_configtxgen_genesis_block {
    command "docker run -it --rm \
    -v $bdir/asset/artifacts/block:/workdir/block \
    -v $bdir/asset/artifacts/crypto-config:/workdir/crypto-config \
    -v $bdir/asset/tool/system-channel-configtx.yaml:/workdir/configtx.yaml \
    --workdir /workdir \
    hyperledger/fabric-tools:$FABRIC_VERSION \
    configtxgen -profile system-channelProfile -channelID system-channel -outputBlock /workdir/block/system-channel.block -configPath /workdir"
}

function blockchain_build_configtxgen_channel_tx {
    command "docker run -it --rm \
    -v $bdir/asset/artifacts/tx:/workdir/tx \
    -v $bdir/asset/artifacts/crypto-config:/workdir/crypto-config \
    -v $bdir/asset/tool/$1-configtx.yaml:/workdir/configtx.yaml \
    --workdir /workdir \
    hyperledger/fabric-tools:$FABRIC_VERSION \
    configtxgen -profile $1Profile -channelID $1 -outputCreateChannelTx /workdir/tx/$1.tx -configPath /workdir"
}

function blockchain_channel_create {
    CHANNEL=$1
    PEER_NAME=$2
    command "docker exec -it \
    cli.$PEER_NAME \
    peer channel create -c $CHANNEL -f /etc/hyperledger/fabric/tx/$CHANNEL.tx --outputBlock /etc/hyperledger/fabric/block/$CHANNEL.block $GLOBAL_FLAGS"
}

function blockchain_channel_join {
    ORG_NAME=$1
    CHANNEL=$2
    command "docker exec -it \
    cli.peer0.$ORG_NAME.example.com \
    peer channel join -b /etc/hyperledger/fabric/block/$CHANNEL.block"
}

function blockchain_chaincode_package {
    PEER_NAME=$1
    CHAINCODE_NAME=$2
    VERSION=$3
    command "docker exec -it \
    cli.$PEER_NAME \
    peer lifecycle chaincode package $CHAINCODE_DIR/$CHAINCODE_NAME-$VERSION.tar.gz --path $CHAINCODE_DIR/$CHAINCODE_NAME --lang golang --label $CHAINCODE_NAME-$VERSION"
}

function blockchain_chaincode_install {
    PEER_NAME=$1
    CHAINCODE_NAME=$2
    VERSION=$3
    command "docker exec -it \
    cli.$PEER_NAME \
    peer lifecycle chaincode install $CHAINCODE_DIR/$CHAINCODE_NAME-$VERSION.tar.gz"
}

function blockchain_chaincode_getpackageid {
    PEER_NAME=$1
    CHAINCODE_NAME=$2
    VERSION=$3
    command "docker exec -it \
    cli.$PEER_NAME \
    peer lifecycle chaincode queryinstalled"
    PACKAGE_ID=$(sed -n "/$CHAINCODE_NAME-$VERSION/{s/^Package ID: //; s/, Label:.*$//; p;}" $bdir/log.txt)
}

function blockchain_chaincode_approveformyorg {
    PEER_NAME=$1
    CHANNEL_NAME=$2
    CHAINCODE_NAME=$3
    VERSION=$4
    SEQUENCE=$5
    blockchain_chaincode_getpackageid $PEER_NAME $CHAINCODE_NAME $VERSION
    command "docker exec -it \
    cli.$PEER_NAME \
    peer lifecycle chaincode approveformyorg \
    --channelID $CHANNEL_NAME \
    --name $CHAINCODE_NAME \
    --version $VERSION \
    --package-id $PACKAGE_ID \
    --collections-config $CHAINCODE_DIR/collections_config.json
    --sequence $SEQUENCE \
    $GLOBAL_FLAGS"
}

function blockchain_chaincode_checkcommitreadiness {
    PEER_NAME=$1
    CHANNEL_NAME=$2
    CHAINCODE_NAME=$3
    VERSION=$4
    SEQUENCE=$5
    command "docker exec -it \
    cli.$PEER_NAME \
    peer lifecycle chaincode checkcommitreadiness \
    --channelID $CHANNEL_NAME \
    --name $CHAINCODE_NAME \
    --version $VERSION \
    --sequence $SEQUENCE \
    $GLOBAL_FLAGS"
}

function blockchain_chaincode_commit {
    PEER_NAME=$1
    CHANNEL_NAME=$2
    CHAINCODE_NAME=$3
    VERSION=$4
    SEQUENCE=$5
    command "docker exec -it \
    cli.$PEER_NAME \
    peer lifecycle chaincode commit \
    --channelID $CHANNEL_NAME \
    --name $CHAINCODE_NAME \
    --version $VERSION \
    --collections-config $CHAINCODE_DIR/collections_config.json
    --sequence $SEQUENCE \
    $GLOBAL_FLAGS"
}

function blockchain_chaincode_querycommitted {
    PEER_NAME=$BTP_PEER
    CHANNEL_NAME=$1
    command "docker exec -it \
    cli.$PEER_NAME \
    peer lifecycle chaincode querycommitted \
    --channelID $CHANNEL_NAME \
    $GLOBAL_FLAGS"
}

function blockchain_chaincode_init {
    PEER_NAME=$BTP_PEER
    CHANNEL_NAME=$1
    CHAINCODE_NAME=$2
    FCN_CALL='{"function":"InitLedger","Args":[]}'
    command "docker exec -it \
    cli.$PEER_NAME \
    peer chaincode invoke \
    --channelID $CHANNEL_NAME \
    --name $CHAINCODE_NAME \
    -c $FCN_CALL \
    $GLOBAL_FLAGS"
}

function blockchain_all {
    blockchain_build
    blockchain_up
    blockchain_channel
    blockchain_chaincode
}

function blockchain_clean {
    blockchain_down
    rm -rf $bdir
}

function blockchain_build {
    mkdir -p $bdir && \
    cp -rf $sdir/asset $bdir/ && \
    cp -rf $sdir/../../../chaincode-go $bdir/asset && \
    mv $bdir/asset/chaincode-go $bdir/asset/chaincodes
    mkdir -p $bdir/asset/artifacts/block && \
    mkdir -p $bdir/asset/artifacts/tx >> /dev/null 2>&1
    blockchain_build_cryptogen
    blockchain_build_configtxgen
}

function blockchain_up {
    docker network create btp >> /dev/null 2>&1
    command "docker-compose -f "$bdir/asset/docker/docker-compose-ca.yaml" up -d"
    command "docker-compose -f "$bdir/asset/docker/docker-compose-peer.yaml" up -d"
    command "docker-compose -f "$bdir/asset/docker/docker-compose-orderer.yaml" up -d"
    cecho "INFO" "Waiting 10s for blockchain network stable"
    sleep 10s
}

function blockchain_down {
    for TARGET in "$bdir/asset/docker"/*
    do
        command "docker-compose -f $TARGET down"
    done
    docker network rm btp >> /dev/null 2>&1
    yes | docker volume prune
}

function blockchain_channel {
    blockchain_channel_create $CHANNEL $BTP_PEER
    sleep 5s
    blockchain_channel_join $ORGANIZATION $CHANNEL
}

function blockchain_chaincode {
    VERSION=1.0
    SEQUENCE=1
    for CHAINCODE in ${CHAINCODES[@]}
    do
        for TARGET in "$bdir/asset/chaincodes"/*
        do
            if [ $TARGET == "$bdir/asset/chaincodes/$CHAINCODE" ]
            then
                blockchain_chaincode_package $BTP_PEER $CHAINCODE $VERSION
                blockchain_chaincode_install $BTP_PEER $CHAINCODE $VERSION
                blockchain_chaincode_approveformyorg $BTP_PEER $CHANNEL $CHAINCODE $VERSION $SEQUENCE
                sleep 1s
                blockchain_chaincode_checkcommitreadiness $BTP_PEER $CHANNEL $CHAINCODE $VERSION $SEQUENCE
                blockchain_chaincode_commit $BTP_PEER $CHANNEL $CHAINCODE $VERSION $SEQUENCE
                blockchain_chaincode_querycommitted $CHANNEL
                blockchain_chaincode_init $CHANNEL $CHAINCODE
            fi
        done
    done
}

function blockchain_chaincode_invoke {
    PEER_NAME=$BTP_PEER
    CHANNEL_NAME=$1
    CHAINCODE_NAME=$2
    QUERY=$3
    command "docker exec -it \
    cli.$PEER_NAME \
    peer chaincode invoke  \
    --channelID $CHANNEL_NAME \
    --name $CHAINCODE_NAME \
    -c $QUERY \
    $GLOBAL_FLAGS"
}

function blockchain_chaincode_query {
    PEER_NAME=$BTP_PEER
    CHANNEL_NAME=$1
    CHAINCODE_NAME=$2
    QUERY=$3
    command "docker exec -it \
    cli.$PEER_NAME \
    peer chaincode query \
    --channelID $CHANNEL_NAME \
    --name $CHAINCODE_NAME \
    -c $QUERY"
}

function blockchain_usage {
    cecho "RED" "not implemantation now"
}

function blockchain_invoke {
    CHANNEL=$1
    CHHAINCODE_NAME=$2
    blockchain_chaincode_invoke $CHANNEL $CHHAINCODE_NAME '{"function":"","Args":[]}'
}

function blockchain_query {
    CHANNEL=$1
    CHHAINCODE_NAME=$2
    blockchain_chaincode_query $CHANNEL $CHHAINCODE_NAME '{"function":"","Args":[]}'
}
function main {
    case $1 in
        all | clean | build | up | down | channel | chaincode | test | invoke | query )
            cmd=blockchain_$1
            shift
            $cmd $@
            ;;
        *)
            blockchain_usage
            exit
            ;;
    esac
}

main $@