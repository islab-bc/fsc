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
func Create(ctx contractapi.TransactionContextInterface, ncr request.NoticeCreateRequest) error {
	if ncr.DID == "" {
		logger.Error("DID가 없음")
		return errors.New("DID가 없음")
	}
	if ncr.Company == "" {
		logger.Error("회사이름이 없음")
		return errors.New("회사이름이 없음")
	}
	if ncr.Writer == "" {
		logger.Error("작성자가 없음")
		return errors.New("작성자가 없음")
	}
	if ncr.Title == "" {
		logger.Error("제목이 없음")
		return errors.New("제목이 없음")
	}
	if ncr.Content == "" {
		logger.Error("본문이 없음")
		return errors.New("본문이 없음")
	}
	if ncr.IsPopUp < 0 || ncr.IsPopUp > 1 {
		logger.Error("팝업 표시 여부 에러")
		return errors.New("팝업 표시 여부 에러")
	}
	if ncr.PopUpContent == "" {
		logger.Error("팝업 내용이 없음")
		return errors.New("팝업 내용이 없음")
	}
	if ncr.EnrolledTime == "" {
		logger.Error("등록시간 에러")
		return errors.New("등록시간 에러")
	}

	noticeModel := model.NewNotice()
	noticeModel.SetDocType("notice")
	noticeModel.SetDID(ncr.DID)
	noticeModel.SetCompany(ncr.Company)
	noticeModel.SetWriter(ncr.Writer)
	noticeModel.SetTitle(ncr.Title)
	noticeModel.SetContent(ncr.Content)
	noticeModel.SetFiles(ncr.Files)
	noticeModel.SetIsPopUp(ncr.IsPopUp)
	noticeModel.SetPopUpContent(ncr.PopUpContent)
	noticeModel.SetEnrolledTime(ncr.EnrolledTime)
	noticeModel.SetUpdatedTime(ncr.EnrolledTime)

	isCreated, err := getNoticeByKey(ctx, noticeModel.GetIFKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("GetState 실패")
	}
	if len(isCreated) != 0 {
		logger.Error("이미 등록된 게시글")
		return errors.New("이미 등록된 게시글")
	}

	noticeModelAsBytes, err := json.Marshal(noticeModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩 실패")
	}

	err = ctx.GetStub().PutState(noticeModel.GetIFKey(), noticeModelAsBytes)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutState 실패")
	}

	return nil
}
