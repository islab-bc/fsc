package contract

import (
	"encoding/json"
	"net/http"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	response "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/response"
	communityService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/community"
	infoshareService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/infoshare"
	fixedNoticeService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/notice/fixed"
	notFixedNoticeService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/notice/notFixed"
	suggestionService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/suggestion"
)

// DeleteFixedNotice is method to delete post for board
func (bc *BoardChaincode) DeleteFixedNotice(ctx contractapi.TransactionContextInterface, rawNoticeDeleteRequest string) *response.ChaincodeResponse {
	var noticeDeleteRequest request.NoticeDeleteRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawNoticeDeleteRequest,
	}

	err := json.Unmarshal([]byte(rawNoticeDeleteRequest), &noticeDeleteRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = fixedNoticeService.Delete(ctx, noticeDeleteRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// DeleteNotFixedNotice is method to delete post for board
func (bc *BoardChaincode) DeleteNotFixedNotice(ctx contractapi.TransactionContextInterface, rawNoticeDeleteRequest string) *response.ChaincodeResponse {
	var noticeDeleteRequest request.NoticeDeleteRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawNoticeDeleteRequest,
	}

	err := json.Unmarshal([]byte(rawNoticeDeleteRequest), &noticeDeleteRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = notFixedNoticeService.Delete(ctx, noticeDeleteRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// DeleteSuggestion is method to delete post for board
func (bc *BoardChaincode) DeleteSuggestion(ctx contractapi.TransactionContextInterface, rawSuggestionDeleteRequest string) *response.ChaincodeResponse {
	var suggestionDeleteRequest request.SuggestionDeleteRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawSuggestionDeleteRequest,
	}

	err := json.Unmarshal([]byte(rawSuggestionDeleteRequest), &suggestionDeleteRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = suggestionService.Delete(ctx, suggestionDeleteRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// DeleteCommunity is method to delete post for board
func (bc *BoardChaincode) DeleteCommunity(ctx contractapi.TransactionContextInterface, rawCommunityDeleteRequest string) *response.ChaincodeResponse {
	var communityDeleteRequest request.CommunityDeleteRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommunityDeleteRequest,
	}

	err := json.Unmarshal([]byte(rawCommunityDeleteRequest), &communityDeleteRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = communityService.Delete(ctx, communityDeleteRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// DeleteInfoshare is method to delete post for board
func (bc *BoardChaincode) DeleteInfoshare(ctx contractapi.TransactionContextInterface, rawInfoshareDeleteRequest string) *response.ChaincodeResponse {
	var infoshareDeleteRequest request.InfoshareDeleteRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawInfoshareDeleteRequest,
	}

	err := json.Unmarshal([]byte(rawInfoshareDeleteRequest), &infoshareDeleteRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = infoshareService.Delete(ctx, infoshareDeleteRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}
