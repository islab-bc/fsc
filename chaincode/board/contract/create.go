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

// CreateFixedNotice is method to create post for board
func (bc *BoardChaincode) CreateFixedNotice(ctx contractapi.TransactionContextInterface, rawNoticeCreateRequest string) *response.ChaincodeResponse {
	var noticeCreateRequest request.NoticeCreateRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawNoticeCreateRequest,
	}

	err := json.Unmarshal([]byte(rawNoticeCreateRequest), &noticeCreateRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = fixedNoticeService.Create(ctx, noticeCreateRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse

}

// CreateNotFixedNotice is method to create post for board
func (bc *BoardChaincode) CreateNotFixedNotice(ctx contractapi.TransactionContextInterface, rawNoticeCreateRequest string) *response.ChaincodeResponse {
	var noticeCreateRequest request.NoticeCreateRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawNoticeCreateRequest,
	}

	err := json.Unmarshal([]byte(rawNoticeCreateRequest), &noticeCreateRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = notFixedNoticeService.Create(ctx, noticeCreateRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse

}

// CreateSuggestion is method to create post for board
func (bc *BoardChaincode) CreateSuggestion(ctx contractapi.TransactionContextInterface, rawSuggestionCreateRequest string) *response.ChaincodeResponse {
	var suggestionCreateRequest request.SuggestionCreateRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawSuggestionCreateRequest,
	}

	err := json.Unmarshal([]byte(rawSuggestionCreateRequest), &suggestionCreateRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = suggestionService.Create(ctx, suggestionCreateRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// CreateCommunity is method to create post for board
func (bc *BoardChaincode) CreateCommunity(ctx contractapi.TransactionContextInterface, rawCommunityCreateRequest string) *response.ChaincodeResponse {
	var communityCreateRequest request.CommunityCreateRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommunityCreateRequest,
	}

	err := json.Unmarshal([]byte(rawCommunityCreateRequest), &communityCreateRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = communityService.Create(ctx, communityCreateRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// CreateInfoshare is method to create post for board
func (bc *BoardChaincode) CreateInfoshare(ctx contractapi.TransactionContextInterface, rawInfoshareCreateRequest string) *response.ChaincodeResponse {
	var infoshareCreateRequest request.InfoshareCreateRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawInfoshareCreateRequest,
	}

	err := json.Unmarshal([]byte(rawInfoshareCreateRequest), &infoshareCreateRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = infoshareService.Create(ctx, infoshareCreateRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}
