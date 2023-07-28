package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/model"
)

// Update is funtion to update post for board
func Update(ctx contractapi.TransactionContextInterface, sur request.SuggestionUpdateRequest) error {
	if sur.Company == "" {
		logger.Error("회사이름이 없음")
		return errors.New("회사이름이 없음")
	}
	if sur.DID == "" {
		logger.Error("DID가 없음")
		return errors.New("DID가 없음")
	}
	if sur.Title == "" {
		logger.Error("제목이 없음")
		return errors.New("제목이 없음")
	}
	if sur.Content == "" {
		logger.Error("본문이 없음")
		return errors.New("본문이 없음")
	}
	if sur.Status != 0 && sur.Status != 1 && sur.Status != 2 {
		logger.Error("상태 에러")
		return errors.New("상태 에러")
	}
	if sur.UpdatedTime == "" {
		logger.Error("수정시간 에러")
		return errors.New("수정시간에러 ")
	}

	targetSuggestionModel := model.NewSuggestion()
	targetSuggestionModel.SetCompany(sur.Company)
	targetSuggestionModel.SetEnrolledTime(sur.EnrolledTime)

	isCreated, err := getSuggestionByKey(ctx, targetSuggestionModel.GetKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("GetState 실패")
	}
	if len(isCreated) == 0 {
		logger.Error("존재하지 않는 게시글")
		return errors.New("존재하지 않는 게시글")
	}

	err = json.Unmarshal([]byte(isCreated), &targetSuggestionModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 디코딩 실패")
	}

	didList, err := InvokeGetDIDHistory(ctx, sur.DID)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("InvokeGetDIDHistory() 호출 실패")
	}
	if CheckPermission(ctx, targetSuggestionModel.GetDID(), didList) == false {
		logger.Error("권한이 없음")
		return errors.New("권한이 없음")
	}

	targetSuggestionModel.SetTitle(sur.Title)
	targetSuggestionModel.SetDID(sur.DID)
	targetSuggestionModel.SetContent(sur.Content)
	targetSuggestionModel.SetFiles(sur.Files)
	targetSuggestionModel.SetStatus(sur.Status)
	targetSuggestionModel.SetUpdatedTime(sur.UpdatedTime)
	targetSuggestionModelAsBytes, err := json.Marshal(targetSuggestionModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩 실패")
	}

	err = ctx.GetStub().PutState(targetSuggestionModel.GetKey(), targetSuggestionModelAsBytes)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutState 실패")
	}

	return nil
}
