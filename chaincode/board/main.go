package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
)

func main() {
	boardChaincode, err := contractapi.NewChaincode(&contract.BoardChaincode{})
	if err != nil {
		logger.Error("Error creating board chaincode")
		panic(err)
	}
	logger.Info("Generate board chaincode")

	if err := boardChaincode.Start(); err != nil {
		logger.Error("Error starting board chaincode:")
		panic(err)
	}
	logger.Info("Strart board chaincode")
}
