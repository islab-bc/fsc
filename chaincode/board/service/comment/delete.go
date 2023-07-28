package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/model"
)

// Delete is function to delete post by key for comment
func Delete(ctx contractapi.TransactionContextInterface, cdr request.CommentDeleteRequest) error {
	if cdr.BoardCompany == "" {
		logger.Error("게시글의 회사이름이 없음")
		return errors.New("게시글의 회사이름이 없음")
	}
	if cdr.DID == "" {
		logger.Error("DID가 없음")
		return errors.New("DID가 없음")
	}
	if cdr.BoardEnrolledTime == "" {
		logger.Error("게시글의 등록시간이 없음")
		return errors.New("게시글의 등록시간이 없음")
	}
	if cdr.Commenter == "" {
		logger.Error("댓글 작성자가 없음")
		return errors.New("댓글 작성자가 없음")
	}
	if cdr.EnrolledTime == "" {
		logger.Error("댓글 등록시간 에러")
		return errors.New("댓글 등록시간 에러")
	}

	commentModel := model.NewComment(
		"comment",
	)
	commentModel.SetBoardCompany(cdr.BoardCompany)
	commentModel.SetBoardEnrolledTime(cdr.BoardEnrolledTime)
	commentModel.SetCommenter(cdr.Commenter)
	commentModel.SetEnrolledTime(cdr.EnrolledTime)

	isCreated, err := getCommentByKey(ctx, commentModel.GetKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("GetState 실패")
	}
	if len(isCreated) == 0 {
		logger.Error("댓글을 찾을 수 없음")
		return errors.New("댓글을 찾을 수 없음")
	}

	err = json.Unmarshal([]byte(isCreated), &commentModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 디코딩 실패")
	}

	didList, err := InvokeGetDIDHistory(ctx, cdr.DID)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("InvokeGetDIDHistory() 호출 실패")
	}
	if CheckPermission(ctx, commentModel.GetDID(), didList) == false {
		logger.Error("권한이 없음")
		return errors.New("권한이 없음")
	}

	err = ctx.GetStub().DelState(commentModel.GetKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("DelState 실패")
	}
	return nil
}
