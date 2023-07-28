package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/model"
)

// Create is function to create board data for board
func Create(ctx contractapi.TransactionContextInterface, icr request.InfoshareCreateRequest) error {
	if icr.DID == "" {
		logger.Error("DID가 없음")
		return errors.New("DID가 없음")
	}
	if icr.Company == "" {
		logger.Error("회사이름이 없음")
		return errors.New("회사이름이 없음")
	}
	if icr.Writer == "" {
		logger.Error("작성자가 없음")
		return errors.New("작성자가 없음")
	}
	if icr.Category == "" {
		logger.Error("말머리가 없음")
		return errors.New("말머리가 없음")
	}
	if icr.Title == "" {
		logger.Error("제목이 없음")
		return errors.New("제목이 없음")
	}
	if icr.Content == "" {
		logger.Error("본문이 없음")
		return errors.New("본문이 없음")
	}
	if icr.EnrolledTime == "" {
		logger.Error("등록시간 에러")
		return errors.New("등록시간 에러")
	}

	infoshareModel := model.NewInfoshare()
	infoshareModel.SetDocType("infoshare")
	infoshareModel.SetDID(icr.DID)
	infoshareModel.SetCompany(icr.Company)
	infoshareModel.SetWriter(icr.Writer)
	infoshareModel.SetCategory(icr.Category)
	infoshareModel.SetTitle(icr.Title)
	infoshareModel.SetContent(icr.Content)
	infoshareModel.SetFiles(icr.Files)
	infoshareModel.SetEnrolledTime(icr.EnrolledTime)
	infoshareModel.SetUpdatedTime(icr.EnrolledTime)

	isCreated, err := getInfoshareByKey(ctx, infoshareModel.GetKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("GetState 실패")
	}
	if len(isCreated) != 0 {
		logger.Error("이미 등록된 게시글임")
		return errors.New("이미 등록된 게시글임")
	}

	infoshareModelAsBytes, err := json.Marshal(infoshareModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩 실패")
	}

	err = ctx.GetStub().PutState(infoshareModel.GetKey(), infoshareModelAsBytes)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutState 실패")
	}

	return nil
}
