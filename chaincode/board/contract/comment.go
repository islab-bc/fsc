package contract

import (
	"encoding/json"
	"net/http"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	response "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/response"
	commentService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/comment"
)

// CreateComment is method to add comment of post for board
func (bc *BoardChaincode) CreateComment(ctx contractapi.TransactionContextInterface, rawCommentCreateRequest string) *response.ChaincodeResponse {
	var commentCreateRequest request.CommentCreateRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommentCreateRequest,
	}

	err := json.Unmarshal([]byte(rawCommentCreateRequest), &commentCreateRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = commentService.Create(ctx, commentCreateRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// UpdateComment is method to update comment of post for board
func (bc *BoardChaincode) UpdateComment(ctx contractapi.TransactionContextInterface, rawCommentUpdateRequest string) *response.ChaincodeResponse {
	var commentUpdateRequest request.CommentUpdateRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommentUpdateRequest,
	}

	err := json.Unmarshal([]byte(rawCommentUpdateRequest), &commentUpdateRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = commentService.Update(ctx, commentUpdateRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// DeleteComment is method to delete comment of post for board
func (bc *BoardChaincode) DeleteComment(ctx contractapi.TransactionContextInterface, rawCommentDeleteRequest string) *response.ChaincodeResponse {
	var commentDeleteRequest request.CommentDeleteRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommentDeleteRequest,
	}

	err := json.Unmarshal([]byte(rawCommentDeleteRequest), &commentDeleteRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = commentService.Delete(ctx, commentDeleteRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// GetCommentsByBoardKey is method to get comments of post by writer with paging for board
func (bc *BoardChaincode) GetCommentsByBoardKey(ctx contractapi.TransactionContextInterface, rawCommentGetByBoardKey string) *response.ChaincodeResponse {
	var commentGetByBoardKeyRequest request.CommentGetByBoardKeyRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommentGetByBoardKey,
	}

	err := json.Unmarshal([]byte(rawCommentGetByBoardKey), &commentGetByBoardKeyRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := commentService.GetCommentsByBoardKeyWithPaging(ctx, commentGetByBoardKeyRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.ChaincodeResult = data
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}
	return &chaincodeResponse
}

// GetCommentsByCommenter is method to get comments of post by writer with paging for board
func (bc *BoardChaincode) GetCommentsByCommenter(ctx contractapi.TransactionContextInterface, rawCommentGetByCommenter string) *response.ChaincodeResponse {
	var commentGetByCommenterRequest request.CommentGetByCommenterRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommentGetByCommenter,
	}

	err := json.Unmarshal([]byte(rawCommentGetByCommenter), &commentGetByCommenterRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := commentService.GetCommentsByCommenterWithPaging(ctx, commentGetByCommenterRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.ChaincodeResult = data
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}
	return &chaincodeResponse
}

// GetCommentsByDID is method to get comments of post by writer with paging for board
func (bc *BoardChaincode) GetCommentsByDID(ctx contractapi.TransactionContextInterface, rawCommentGetByDID string) *response.ChaincodeResponse {
	var commentGetByDIDRequest request.CommentGetByDIDRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommentGetByDID,
	}

	err := json.Unmarshal([]byte(rawCommentGetByDID), &commentGetByDIDRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := commentService.GetCommentListByDIDWithPaging(ctx, commentGetByDIDRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.ChaincodeResult = data
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}
	return &chaincodeResponse
}
