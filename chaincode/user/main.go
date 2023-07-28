package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract"
)

func main() {
	userChaincode, err := contractapi.NewChaincode(&contract.UserChaincode{})
	if err != nil {
		log.Panicf("체인코드 생성에 실패함: %v", err)
	}

	if err := userChaincode.Start(); err != nil {
		log.Panicf("체인코드 생성에 실패함: %v", err)
	}
}
