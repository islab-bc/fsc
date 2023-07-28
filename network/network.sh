#!/bin/bash

function network_start {
    sudo bash ./blockchain/blockchain.sh all
}

function network_invoke {
    CHANNEL=$1
    CHAINCODE_NAME=$2
    sudo bash ./blockchain/blockchain.sh invoke $CHANNEL $CHAINCODE_NAME
}

function network_query {
    CHANNEL=$1
    CHAINCODE_NAME=$2
    sudo bash ./blockchain/blockchain.sh query $CHANNEL $CHAINCODE_NAME
}

function network_clear {
    docker rm -f -v `docker ps -aqf "name=ca.btp.example.com"`
    docker rm -f -v `docker ps -aqf "name=cli.peer0.btp.example.com"`
    docker rm -f -v `docker ps -aqf "name=couchdb.peer0.btp.example.com"`
    docker rm -f -v `docker ps -aqf "name=btp-peer0.btp.example.com/*"`
    docker rm -f -v `docker ps -aqf "name=orderer0.example.com"`
    docker rm -f -v `docker ps -aqf "name=peer0.btp.example.com"`
    sleep 3s
    sudo bash ./blockchain/blockchain.sh clean
}

function network_restart {
    network_clear
    sleep 1s
    network_start
}

function network_usage {
    cecho "RED" "you should use fabric | clear | invoke | query | restart"
}

function main {
    case $1 in
        start | clear | invoke | query | restart)
            cmd=network_$1
            shift
            $cmd $@
            ;;
        *)
            network_usage
            exit
            ;;
    esac
}

main $@