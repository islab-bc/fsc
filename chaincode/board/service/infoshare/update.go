package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/model"
)

// Update is funtion to update post for infoshare board
func Update(ctx contractapi.TransactionContextInterface, iur request.InfoshareUpdateRequest) error {
	if iur.Company == "" {
		logger.Error("회사이름이 없음")
		return errors.New("회사이름이 없음")
	}
	if iur.DID == "" {
		logger.Error("DID가 없음")
		return errors.New("DID가 없음")
	}
	if iur.Category == "" {
		logger.Error("말머리가 없음")
		return errors.New("말머리가 없음")
	}
	if iur.Title == "" {
		logger.Error("제목이 없음")
		return errors.New("제목이 없음")
	}
	if iur.Content == "" {
		logger.Error("본문이 없음")
		return errors.New("본문이 없음")
	}
	if iur.EnrolledTime == "" {
		logger.Error("등록시간 에러")
		return errors.New("등록시간 에러")
	}
	if iur.UpdatedTime == "" {
		logger.Error("수정시간 에러")
		return errors.New("수정시간에러 ")
	}

	targetInfoshareModel := model.NewInfoshare()
	targetInfoshareModel.SetCompany(iur.Company)
	targetInfoshareModel.SetEnrolledTime(iur.EnrolledTime)

	isCreated, err := getInfoshareByKey(ctx, targetInfoshareModel.GetKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("GetState 실패")
	}
	if len(isCreated) == 0 {
		logger.Error("존재하지 않는 게시글")
		return errors.New("존재하지 않는 게시글")
	}

	err = json.Unmarshal([]byte(isCreated), &targetInfoshareModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 디코딩 실패")
	}
	didList, err := InvokeGetDIDHistory(ctx, iur.DID)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("InvokeGetDIDHistory() 호출 실패")
	}
	if CheckPermission(ctx, targetInfoshareModel.GetDID(), didList) == false {
		logger.Error("권한이 없음")
		return errors.New("권한이 없음")
	}

	targetInfoshareModel.SetCategory(iur.Category)
	targetInfoshareModel.SetTitle(iur.Title)
	targetInfoshareModel.SetDID(iur.DID)
	targetInfoshareModel.SetContent(iur.Content)
	targetInfoshareModel.SetFiles(iur.Files)
	targetInfoshareModel.SetUpdatedTime(iur.UpdatedTime)
	targetInfoshareModelAsBytes, err := json.Marshal(targetInfoshareModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩 실패")
	}

	err = ctx.GetStub().PutState(targetInfoshareModel.GetKey(), targetInfoshareModelAsBytes)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutState 실패")
	}

	return nil
}
