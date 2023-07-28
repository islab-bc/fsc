package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/model"
)

// Enroll ...
func Enroll(ctx contractapi.TransactionContextInterface, aer request.AdmissionEnrollRequest) error {
	if aer.DID == "" {
		logger.Error("유효하지 않은 DID")
    	return errors.New("유효하지 않은 DID")
	}
	if aer.Location == "" {
		logger.Error("유효하지 않은 Location")
    	return errors.New("유효하지 않은 Location")
	}
	if aer.EnrolledTime == "" {
		logger.Error("출입 시간이 입력되지 않음")
    	return errors.New("출입 시간이 입력되지 않음")
	}
	admissionModel := model.NewAdmission()
	admissionModel.SetDID(aer.DID)
	admissionModel.SetLocation(aer.Location)
	admissionModel.SetEnrolledTime(aer.EnrolledTime)

	result, err := json.Marshal(admissionModel)
	if err != nil {
		logger.Error(err.Error())
    	return errors.New("JSON 인코딩에 실패함")
	}

	err = ctx.GetStub().PutPrivateData("collectionUsers", admissionModel.GetKey(), result)
	if err != nil {
		logger.Error(err.Error())
    	return errors.New("PutPrivateData() 에러 발생")
	}

	return nil
}