package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/model"
)

// Create is funtion to add comment of post for board
func Create(ctx contractapi.TransactionContextInterface, ccr request.CommentCreateRequest) error {
	if ccr.DID == "" {
		logger.Error("DID가 없음")
		return errors.New("DID가 없음")
	}
	if ccr.BoardCompany == "" {
		logger.Error("게시글의 회사이름이 없음")
		return errors.New("게시글의 회사이름이 없음")
	}
	if ccr.BoardEnrolledTime == "" {
		logger.Error("게시글의 등록시간이 없음")
		return errors.New("게시글의 등록시간이 없음")
	}
	if ccr.Commenter == "" {
		logger.Error("댓글 작성자가 없음")
		return errors.New("댓글 작성자가 없음")
	}
	if ccr.Content == "" {
		logger.Error("댓글 내용이 없음")
		return errors.New("댓글 내용이 없음")
	}
	if ccr.EnrolledTime == "" {
		logger.Error("댓글 등록시간 에러")
		return errors.New("댓글 등록시간 에러")
	}
	commentModel := model.NewComment(
		"comment",
	)
	commentModel.SetDID(ccr.DID)
	commentModel.SetBoardCompany(ccr.BoardCompany)
	commentModel.SetBoardEnrolledTime(ccr.BoardEnrolledTime)
	commentModel.SetParentCommenter(ccr.ParentCommenter)
	commentModel.SetParentEnrolledTime(ccr.ParentEnrolledTime)
	commentModel.SetCommenter(ccr.Commenter)
	commentModel.SetContent(ccr.Content)
	commentModel.SetEnrolledTime(ccr.EnrolledTime)
	commentModel.SetUpdatedTime(ccr.EnrolledTime)
	commentModel.SetHierarchyIndex()

	boardKey := commentModel.GetBoardKey()
	if isBoardKeyAvailable(ctx, boardKey) == false {
		logger.Error("게시글을 찾을 수 없음")
		return errors.New("게시글을 찾을 수 없음")
	}

	isUploaded, err := getCommentByKey(ctx, commentModel.GetKey())
	if err != nil {
		logger.Error(err.Error())
		return errors.New("GetState 실패")
	}
	if len(isUploaded) != 0 {
		logger.Error("이미 등록된 댓글임")
		return errors.New("이미 등록된 댓글임")
	}

	commentModelAsBytes, err := json.Marshal(commentModel)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("JSON 인코딩 실패")
	}

	err = ctx.GetStub().PutState(commentModel.GetKey(), commentModelAsBytes)
	if err != nil {
		logger.Error(err.Error())
		return errors.New("Putstate 실패")
	}

	return nil
}

// isBoardKeyAvailable is funtion to add comment of post for board
func isBoardKeyAvailable(ctx contractapi.TransactionContextInterface, boardKey string) bool {
	boardIFKey := "IF_" + boardKey
	boardINFKey := "INF_" + boardKey
	boardModel, err := getBoardByKey(ctx, boardKey)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	boardIFModel, err := getBoardByKey(ctx, boardIFKey)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	boardINFModel, err := getBoardByKey(ctx, boardINFKey)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	if len(boardModel) == 0 && len(boardIFModel) == 0 && len(boardINFModel) == 0 {
		logger.Error("게시글을 찾을 수 없음")
		return false
	}

	return true
}
