package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/model"
)

// Update is funtion to update post for comment
func Update(ctx contractapi.TransactionContextInterface, cur request.CommentUpdateRequest) error {
	if cur.BoardCompany == "" {
		logger.Error("게시글의 회사이름이 없음")
		return errors.New("게시글의 회사이름이 없음")
	}
	if cur.DID == "" {
		logger.Error("DID가 없음")
		return errors.New("DID가 없음")
	}
	if cur.BoardEnrolledTime == "" {
		logger.Error("게시글의 등록시간이 없음")
		return errors.New("게시글의 등록시간이 없음")
	}
	if cur.Commenter == "" {
		logger.Error("댓글 작성자가 없음")
		return errors.New("댓글 작성자가 없음")
	}
	if cur.Content == "" {
		logger.Error("댓글 내용이 없음")
		return errors.New("댓글 내용이 없음")
	}
	if cur.EnrolledTime == "" {
		logger.Error("댓글 등록시간 에러")
		return errors.New("댓글 등록시간 에러")
	}
	if cur.UpdatedTime == "" {
		logger.Error("댓글 수정시간 에러")
		return errors.New("댓글 수정시간 에러")
	}

	targetCommentModel := model.NewComment(
		"comment",
	)
	targetCommentModel.SetBoardCompany(cur.BoardCompany)
	targetCommentModel.SetBoardEnrolledTime(cur.BoardEnrolledTime)
	targetCommentModel.SetCommenter(cur.Commenter)
	targetCommentModel.SetEnrolledTime(cur.EnrolledTime)

	isCreated, err := getCommentByKey(ctx, targetCommentModel.GetKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("GetState 실패")
	}
	if len(isCreated) == 0 {
		logger.Error("댓글을 찾을 수 없음")
		return errors.New("댓글을 찾을 수 없음")
	}

	err = json.Unmarshal([]byte(isCreated), &targetCommentModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 디코딩 실패")
	}

	didList, err := InvokeGetDIDHistory(ctx, cur.DID)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("InvokeGetDIDHistory() 호출 실패")
	}

	if CheckPermission(ctx, targetCommentModel.GetDID(), didList) == false {
		logger.Error("권한이 없음")
		return errors.New("권한이 없음")
	}

	targetCommentModel.SetContent(cur.Content)
	targetCommentModel.SetDID(cur.DID)
	targetCommentModel.SetUpdatedTime(cur.UpdatedTime)

	targetCommentModelAsBytes, err := json.Marshal(targetCommentModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩 실패")
	}

	err = ctx.GetStub().PutState(targetCommentModel.GetKey(), targetCommentModelAsBytes)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("PutState 실패")
	}

	return nil
}
