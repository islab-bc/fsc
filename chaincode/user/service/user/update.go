package service

import (
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/model"
)

// UpdateDID ...
func UpdateDID(ctx contractapi.TransactionContextInterface, uur request.UserDIDUpdateRequest) error {
	if uur.DID == "" {
		logger.Error("유효하지 않은 DID")
    	return errors.New("유효하지 않은 DID")
	}
	userModel := model.NewUser()
	userModel.SetName(uur.Name)
	userModel.SetPhoneNumber(uur.PhoneNumber)
	userModel.SetDeviceToken(uur.DeviceToken)
	userModel.SetMobileOS(uur.MobileOS)
	userModel.SetCI(uur.CI)
	userModel.SetImagePath(uur.ImagePath)
	userModel.SetCompanyName(uur.CompanyName)
	userModel.SetDepartment(uur.Department)
	userModel.SetPosition(uur.Position)
	userModel.SetBirthday(uur.Birthday)
	userModel.SetState(uur.State)
	userModel.SetEnrolledTime(uur.EnrolledTime)

	err := ctx.GetStub().PutState(uur.DID, []byte(uur.BeforeDID))
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutState() 에러 발생")
	}

	err = ctx.GetStub().DelPrivateData("collectionUsers", uur.BeforeDID)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("DelPrivateData() 에러 발생")
	}

	result, err := json.Marshal(userModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩에 실패함")
	}

	err = ctx.GetStub().PutPrivateData("collectionUsers", uur.DID, result)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutPrivateData() 에러 발생")
	}

	return nil
}

// UpdateUser ...
func UpdateUser(ctx contractapi.TransactionContextInterface, uur request.UserDIDUpdateRequest) error {
	if uur.DID == "" {
		logger.Error("유효하지 않은 DID")
    	return errors.New("유효하지 않은 DID")
	}
	if uur.Name == "" {
		logger.Error("이름 입력 오류 발생")
    	return errors.New("이름 입력 오류 발생")
	}
	if uur.PhoneNumber == "" {
		logger.Error("연락처 입력 오류 발생")
    	return errors.New("연락처 입력 오류 발생")
	}
	if uur.CompanyName == "" {
		logger.Error("회사명 입력 오류 발생")
    	return errors.New("회사명 입력 오류 발생")
	}
	if uur.Department == "" {
		logger.Error("부서명 입력 오류 발생")
    	return errors.New("부서명 입력 오류 발생")
	}
	if uur.Position == "" {
		logger.Error("직책 입력 오류 발생")
    	return errors.New("직책 입력 오류 발생")
	}

	userModel := model.NewUser()
	userModel.SetName(uur.Name)
	userModel.SetPhoneNumber(uur.PhoneNumber)
	userModel.SetDeviceToken(uur.DeviceToken)
	userModel.SetMobileOS(uur.MobileOS)
	userModel.SetCI(uur.CI)
	userModel.SetImagePath(uur.ImagePath)
	userModel.SetCompanyName(uur.CompanyName)
	userModel.SetDepartment(uur.Department)
	userModel.SetPosition(uur.Position)
	userModel.SetBirthday(uur.Birthday)
	userModel.SetState(uur.State)
	userModel.SetEnrolledTime(uur.EnrolledTime)

	err := ctx.GetStub().DelPrivateData("collectionUsers", uur.BeforeDID)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("DelPrivateData() 에러 발생")
	}

	result, err := json.Marshal(userModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩에 실패함")
	}

	err = ctx.GetStub().PutPrivateData("collectionUsers", uur.DID, result)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutPrivateData() 에러 발생")
	}

	return nil
}

// UpdateState ...
func UpdateState(ctx contractapi.TransactionContextInterface, uur request.UserDIDUpdateRequest) error {
	if uur.DID == "" {
		logger.Error("유효하지 않은 DID")
    	return errors.New("유효하지 않은 DID")
	}
	if uur.State < 0 || uur.State > 4 {
		logger.Error("잘못된 State 입력 오류 발생")
    	return errors.New("잘못된 State 입력 오류 발생")
	}
	userModel := model.NewUser()
	userModel.SetName(uur.Name)
	userModel.SetPhoneNumber(uur.PhoneNumber)
	userModel.SetDeviceToken(uur.DeviceToken)
	userModel.SetMobileOS(uur.MobileOS)
	userModel.SetCI(uur.CI)
	userModel.SetImagePath(uur.ImagePath)
	userModel.SetCompanyName(uur.CompanyName)
	userModel.SetDepartment(uur.Department)
	userModel.SetPosition(uur.Position)
	userModel.SetBirthday(uur.Birthday)
	userModel.SetState(uur.State)
	userModel.SetEnrolledTime(uur.EnrolledTime)

	err := ctx.GetStub().DelPrivateData("collectionUsers", uur.BeforeDID)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("DelPrivateData() 에러 발생")
	}

	result, err := json.Marshal(userModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩에 실패함")
	}

	err = ctx.GetStub().PutPrivateData("collectionUsers", uur.DID, result)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutPrivateData() 에러 발생")
	}

	return nil
}