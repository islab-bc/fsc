package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/model"
)

// Update is funtion to update post for community board
func Update(ctx contractapi.TransactionContextInterface, cur request.CommunityUpdateRequest) error {
	if cur.Company == "" {
		logger.Error("회사이름이 없음")
		return errors.New("회사이름이 없음")
	}
	if cur.DID == "" {
		logger.Error("DID가 없음")
		return errors.New("DID가 없음")
	}
	if cur.Category == "" {
		logger.Error("말머리가 없음")
		return errors.New("말머리가 없음")
	}
	if cur.Title == "" {
		logger.Error("제목이 없음")
		return errors.New("제목이 없음")
	}
	if cur.Content == "" {
		logger.Error("본문이 없음")
		return errors.New("본문이 없음")
	}
	if cur.EnrolledTime == "" {
		logger.Error("등록시간 에러")
		return errors.New("등록시간 에러")
	}
	if cur.UpdatedTime == "" {
		logger.Error("수정시간 에러")
		return errors.New("수정시간에러 ")
	}

	targetCommunityModel := model.NewCommunity()
	targetCommunityModel.SetCompany(cur.Company)
	targetCommunityModel.SetEnrolledTime(cur.EnrolledTime)

	isCreated, err := getCommunityByKey(ctx, targetCommunityModel.GetKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("GetState 실패")
	}
	if len(isCreated) == 0 {
		logger.Error("존재하지 않는 게시글")
		return errors.New("존재하지 않는 게시글")
	}

	err = json.Unmarshal([]byte(isCreated), &targetCommunityModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 디코딩 실패")
	}

	didList, err := InvokeGetDIDHistory(ctx, cur.DID)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("InvokeGetDIDHistory() 호출 실패")
	}
	if CheckPermission(ctx, targetCommunityModel.GetDID(), didList) == false {
		logger.Error("권한이 없음")
		return errors.New("권한이 없음")
	}

	targetCommunityModel.SetCategory(cur.Category)
	targetCommunityModel.SetDID(cur.DID)
	targetCommunityModel.SetTitle(cur.Title)
	targetCommunityModel.SetContent(cur.Content)
	targetCommunityModel.SetFiles(cur.Files)
	targetCommunityModel.SetUpdatedTime(cur.UpdatedTime)
	targetCommunityModelAsBytes, err := json.Marshal(targetCommunityModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩 실패")
	}

	err = ctx.GetStub().PutState(targetCommunityModel.GetKey(), targetCommunityModelAsBytes)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutState 실패")
	}

	return nil
}
