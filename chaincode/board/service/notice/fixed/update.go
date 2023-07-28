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
func Update(ctx contractapi.TransactionContextInterface, nur request.NoticeUpdateRequest) error {
	if nur.Company == "" {
		logger.Error("회사이름이 없음")
		return errors.New("회사이름이 없음")
	}
	if nur.DID == "" {
		logger.Error("DID가 없음")
		return errors.New("DID가 없음")
	}
	if nur.Title == "" {
		logger.Error("제목이 없음")
		return errors.New("제목이 없음")
	}
	if nur.Content == "" {
		logger.Error("본문이 없음")
		return errors.New("본문이 없음")
	}
	if nur.IsPopUp < 0 || nur.IsPopUp > 1 {
		logger.Error("팝업 표시 여부 에러")
		return errors.New("팝업 표시 여부 에러")
	}
	if nur.PopUpContent == "" {
		logger.Error("팝업 내용이 없음")
		return errors.New("팝업 내용이 없음")
	}
	if nur.EnrolledTime == "" {
		logger.Error("등록시간 에러")
		return errors.New("등록시간 에러")
	}
	if nur.UpdatedTime == "" {
		logger.Error("수정시간 에러")
		return errors.New("수정시간에러 ")
	}

	targetNoticeModel := model.NewNotice()
	targetNoticeModel.SetDocType("notice")
	targetNoticeModel.SetCompany(nur.Company)
	targetNoticeModel.SetEnrolledTime(nur.EnrolledTime)
	targetNoticeModel.SetTitle(nur.Title)
	targetNoticeModel.SetContent(nur.Content)
	targetNoticeModel.SetDID(nur.DID)
	targetNoticeModel.SetFiles(nur.Files)
	targetNoticeModel.SetIsPopUp(nur.IsPopUp)
	targetNoticeModel.SetPopUpContent(nur.PopUpContent)
	targetNoticeModel.SetUpdatedTime(nur.UpdatedTime)

	didList, err := invokeGetDIDHistory(ctx, nur.DID)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("InvokeGetDIDHistory() 호출 실패")
	}
	if checkPermission(ctx, targetNoticeModel.GetDID(), didList) == false {
		logger.Error("권한이 없음")
		return errors.New("권한이 없음")
	}

	IFCreated, err := getNoticeByKey(ctx, targetNoticeModel.GetIFKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("IFKey GetState 실패")
	}
	INFCreated, err := getNoticeByKey(ctx, targetNoticeModel.GetINFKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("INFKey GetState 실패")
	}
	if len(IFCreated) == 0 && len(INFCreated) == 0 { // 둘다 존재 하지 않을경우 업데이트 실패
		logger.Error("존재하지 않는 게시글")
		return errors.New("존재하지 않는 게시글")
	}
	if len(IFCreated) != 0 && len(INFCreated) != 0 { //둘다 존재할경우 에러 처리
		logger.Error("게시글이 FIXED, NOTFIXED 함께 존재함")
		return errors.New("게시글이 FIXED, NOTFIXED 함께 존재함")
	}
	if len(IFCreated) == 0 { // not fixed만 존재할 경우 기존값 삭제 후 fixed로 업데이트
		err = updateWithDelete(ctx, targetNoticeModel)
		if err != nil {
			logger.Error(err.Error())
			return errors.New("updateWithDelete 실패")
		}
	}
	if len(INFCreated) == 0 { // fixed만 존재하는 경우 기존값에 덮어쓰기
		err = overwriteOnly(ctx, targetNoticeModel)
		if err != nil {
			logger.Error(err.Error())
			return errors.New("updateOnly 실패")
		}
	}

	return nil
}

func overwriteOnly(ctx contractapi.TransactionContextInterface, targetNoticeModel *model.Notice) error {
	targetNoticeModelAsBytes, err := json.Marshal(targetNoticeModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩 실패")
	}
	err = ctx.GetStub().PutState(targetNoticeModel.GetIFKey(), targetNoticeModelAsBytes) // overwrite with IFKey
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutState 실패")
	}

	return nil
}

func updateWithDelete(ctx contractapi.TransactionContextInterface, targetNoticeModel *model.Notice) error {
	targetNoticeModelAsBytes, err := json.Marshal(targetNoticeModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩 실패")
	}
	err = ctx.GetStub().DelState(targetNoticeModel.GetINFKey()) // delete INF Key
	if err != nil {
		logger.Error(err.Error())
		return errors.New("DelState() 에러 발생")
	}
	err = ctx.GetStub().PutState(targetNoticeModel.GetIFKey(), targetNoticeModelAsBytes) // put IFKey
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutState 실패")
	}

	return nil
}
