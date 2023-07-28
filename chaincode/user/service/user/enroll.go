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
func Enroll(ctx contractapi.TransactionContextInterface, uer request.UserEnrollRequest) error {
	if uer.DID == "" {
		logger.Error("유효하지 않은 DID")
    	return errors.New("유효하지 않은 DID")
	}
	if uer.Name == "" {
		logger.Error("이름이 입력되지 않음")
    	return errors.New("이름이 입력되지 않음")
	}
	if uer.PhoneNumber == "" {
		logger.Error("전화번호가 입력되지 않음")
    	return errors.New(" 입력되지 않음")
	}
	if uer.CI == "" {
		logger.Error("CI이 입력되지 않음")
    	return errors.New("CI이 입력되지 않음")
	}
	if uer.CompanyName == "" {
		logger.Error("회사명이 입력되지 않음")
    	return errors.New("회사명이 입력되지 않음")
	}
	if uer.Department == "" {
		logger.Error("부서가 입력되지 않음")
    	return errors.New("부서가 입력되지 않음")
	}
	if uer.Position == "" {
		logger.Error("직급이 입력되지 않음")
    	return errors.New("직급이 입력되지 않음")
	}
	if uer.Birthday == "" {
		logger.Error("생일이 입력되지 않음")
    	return errors.New("생일이 입력되지 않음")
	}
	if uer.EnrolledTime == "" {
		logger.Error("등록시간이 입력되지 않음")
    	return errors.New("등록시간이 입력되지 않음")
	}
	userModel := model.NewUser()
	userModel.SetName(uer.Name)
	userModel.SetPhoneNumber(uer.PhoneNumber)
	userModel.SetDeviceToken(uer.DeviceToken)
	userModel.SetMobileOS(uer.MobileOS)
	userModel.SetCI(uer.CI)
	userModel.SetImagePath(uer.ImagePath)
	userModel.SetCompanyName(uer.CompanyName)
	userModel.SetDepartment(uer.Department)
	userModel.SetPosition(uer.Position)
	userModel.SetBirthday(uer.Birthday)
	userModel.SetState(uer.State)
	userModel.SetEnrolledTime(uer.EnrolledTime)

	result, err := json.Marshal(userModel)
	if err != nil {
		logger.Error(err.Error())
    	return errors.New("JSON 인코딩에 실패함")
	}

	err = ctx.GetStub().PutPrivateData("collectionUsers", uer.DID, result)
	if err != nil {
		logger.Error(err.Error())
    	return errors.New("PutPrivateData() 에러 발생")
	}

	return nil
}