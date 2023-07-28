package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/model"
)

// Create is function to create suggetsion board
func Create(ctx contractapi.TransactionContextInterface, scr request.SuggestionCreateRequest) error {
	if scr.DID == "" {
		logger.Error("DID가 없음")
		return errors.New("DID가 없음")
	}
	if scr.Company == "" {
		logger.Error("회사이름이 없음")
		return errors.New("회사이름이 없음")
	}
	if scr.Writer == "" {
		logger.Error("작성자가 없음")
		return errors.New("작성자가 없음")
	}
	if scr.Title == "" {
		logger.Error("제목이 없음")
		return errors.New("제목이 없음")
	}
	if scr.Content == "" {
		logger.Error("본문이 없음")
		return errors.New("본문이 없음")
	}
	if scr.Status != 0 && scr.Status != 1 && scr.Status != 2 {
		logger.Error("상태 에러")
		return errors.New("상태 에러")
	}
	if scr.EnrolledTime == "" {
		logger.Error("등록시간 에러")
		return errors.New("등록시간 에러")
	}

	suggestionModel := model.NewSuggestion()
	suggestionModel.SetDocType("suggestion")
	suggestionModel.SetDID(scr.DID)
	suggestionModel.SetCompany(scr.Company)
	suggestionModel.SetWriter(scr.Writer)
	suggestionModel.SetTitle(scr.Title)
	suggestionModel.SetContent(scr.Content)
	suggestionModel.SetFiles(scr.Files)
	suggestionModel.SetStatus(scr.Status)
	suggestionModel.SetEnrolledTime(scr.EnrolledTime)
	suggestionModel.SetUpdatedTime(scr.EnrolledTime)

	isCreated, err := getSuggestionByKey(ctx, suggestionModel.GetKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("GetState 실패")
	}
	if len(isCreated) != 0 {
		logger.Error("이미 등록된 게시글임")
		return errors.New("이미 등록된 게시글임")
	}

	suggestionModelAsBytes, err := json.Marshal(suggestionModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩 실패")
	}

	err = ctx.GetStub().PutState(suggestionModel.GetKey(), suggestionModelAsBytes)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutState 실패")
	}

	return nil
}
