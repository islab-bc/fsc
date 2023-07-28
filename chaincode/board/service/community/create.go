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
func Create(ctx contractapi.TransactionContextInterface, ccr request.CommunityCreateRequest) error {
	if ccr.DID == "" {
		logger.Error("DID가 없음")
		return errors.New("DID가 없음")
	}
	if ccr.Company == "" {
		logger.Error("회사이름이 없음")
		return errors.New("회사이름이 없음")
	}
	if ccr.Writer == "" {
		logger.Error("작성자가 없음")
		return errors.New("작성자가 없음")
	}
	if ccr.Category == "" {
		logger.Error("말머리가 없음")
		return errors.New("말머리가 없음")
	}
	if ccr.Title == "" {
		logger.Error("제목이 없음")
		return errors.New("제목이 없음")
	}
	if ccr.Content == "" {
		logger.Error("본문이 없음")
		return errors.New("본문이 없음")
	}
	if ccr.EnrolledTime == "" {
		logger.Error("등록시간 에러")
		return errors.New("등록시간 에러")
	}

	communityModel := model.NewCommunity()
	communityModel.SetDocType("community")
	communityModel.SetDID(ccr.DID)
	communityModel.SetCompany(ccr.Company)
	communityModel.SetWriter(ccr.Writer)
	communityModel.SetCategory(ccr.Category)
	communityModel.SetTitle(ccr.Title)
	communityModel.SetContent(ccr.Content)
	communityModel.SetFiles(ccr.Files)
	communityModel.SetEnrolledTime(ccr.EnrolledTime)
	communityModel.SetUpdatedTime(ccr.EnrolledTime)

	isCreated, err := getCommunityByKey(ctx, communityModel.GetKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("GetState 실패")
	}
	if len(isCreated) != 0 {
		logger.Error("이미 등록된 게시글")
		return errors.New("이미 등록된 게시글")
	}

	communityModelAsBytes, err := json.Marshal(communityModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩 실패")
	}

	err = ctx.GetStub().PutState(communityModel.GetKey(), communityModelAsBytes)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutState 실패")
	}

	return nil
}
