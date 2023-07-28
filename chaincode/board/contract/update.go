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

// UpdateFixedNotice is method to update post for board
func (bc *BoardChaincode) UpdateFixedNotice(ctx contractapi.TransactionContextInterface, rawNoticeUpdateRequest string) *response.ChaincodeResponse {
	var noticeUpdateRequest request.NoticeUpdateRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawNoticeUpdateRequest,
	}

	err := json.Unmarshal([]byte(rawNoticeUpdateRequest), &noticeUpdateRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = fixedNoticeService.Update(ctx, noticeUpdateRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// UpdateNotFixedNotice is method to update post for board
func (bc *BoardChaincode) UpdateNotFixedNotice(ctx contractapi.TransactionContextInterface, rawNoticeUpdateRequest string) *response.ChaincodeResponse {
	var noticeUpdateRequest request.NoticeUpdateRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawNoticeUpdateRequest,
	}

	err := json.Unmarshal([]byte(rawNoticeUpdateRequest), &noticeUpdateRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = notFixedNoticeService.Update(ctx, noticeUpdateRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// UpdateSuggestion is method to update post for board
func (bc *BoardChaincode) UpdateSuggestion(ctx contractapi.TransactionContextInterface, rawSuggestionUpdateRequest string) *response.ChaincodeResponse {
	var suggestionUpdateRequest request.SuggestionUpdateRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawSuggestionUpdateRequest,
	}

	err := json.Unmarshal([]byte(rawSuggestionUpdateRequest), &suggestionUpdateRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = suggestionService.Update(ctx, suggestionUpdateRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// UpdateCommunity is method to update post for board
func (bc *BoardChaincode) UpdateCommunity(ctx contractapi.TransactionContextInterface, rawCommunityUpdateRequest string) *response.ChaincodeResponse {
	var communityUpdateRequest request.CommunityUpdateRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommunityUpdateRequest,
	}

	err := json.Unmarshal([]byte(rawCommunityUpdateRequest), &communityUpdateRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = communityService.Update(ctx, communityUpdateRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// UpdateInfoshare is method to update post for board
func (bc *BoardChaincode) UpdateInfoshare(ctx contractapi.TransactionContextInterface, rawInfoshareUpdateRequest string) *response.ChaincodeResponse {
	var infoshareUpdateRequest request.InfoshareUpdateRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawInfoshareUpdateRequest,
	}

	err := json.Unmarshal([]byte(rawInfoshareUpdateRequest), &infoshareUpdateRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = infoshareService.Update(ctx, infoshareUpdateRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}
