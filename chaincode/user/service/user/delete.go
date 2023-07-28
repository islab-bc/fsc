package service

import (
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/logger"
)

// Delete ...
func Delete(ctx contractapi.TransactionContextInterface, udr request.UserDeleteRequest) error {
	if udr.DID == "" {
		logger.Error("유효하지 않은 DID")
		return errors.New("유효하지 않은 DID")
	}
	isEnrolled, err := getUserByKey(ctx, udr.DID)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("사용자 정보를 찾을 수 없음")
	}
	if len(isEnrolled) == 0 {
		logger.Error("해당 사용자를 찾을 수 없음")
		return errors.New("해당 사용자를 찾을 수 없음")
	}

	err = ctx.GetStub().DelPrivateData("collectionUsers", udr.DID)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("DelPrivateData() 에러 발생")
	}

	return nil
}
