package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/model"
)

// Delete is function to delete post by key for notice
func Delete(ctx contractapi.TransactionContextInterface, ndr request.NoticeDeleteRequest) error {
	if ndr.Company == "" {
		logger.Error("회사이름이 없음")
		return errors.New("회사이름이 없음")
	}
	if ndr.EnrolledTime == "" {
		logger.Error("등록시간 에러")
		return errors.New("등록시간 에러")
	}
	if ndr.DID == "" {
		logger.Error("DID가 없음")
		return errors.New("DID가 없음")
	}

	noticeModel := model.NewNotice()
	noticeModel.SetCompany(ndr.Company)
	noticeModel.SetEnrolledTime(ndr.EnrolledTime)

	isCreated, err := getNoticeByKey(ctx, noticeModel.GetIFKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("GetState 실패")
	}
	if len(isCreated) == 0 {
		logger.Error("게시글을 찾을 수 없음")
		return errors.New("게시글을 찾을 수 없음")
	}

	err = json.Unmarshal([]byte(isCreated), &noticeModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 디코딩 실패")
	}

	didList, err := invokeGetDIDHistory(ctx, ndr.DID)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("InvokeGetDIDHistory() 호출 실패")
	}
	if checkPermission(ctx, noticeModel.GetDID(), didList) == false {
		logger.Error("권한이 없음")
		return errors.New("권한이 없음")
	}

	err = ctx.GetStub().DelState(noticeModel.GetIFKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("DelState 실패")
	}
	return nil
}
