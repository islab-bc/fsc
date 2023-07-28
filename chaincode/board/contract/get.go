package contract

import (
	"encoding/json"
	"net/http"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	response "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/response"
	boardService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/board"
	communityService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/community"
	infoshareService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/infoshare"
	commonNoticeService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/notice/common"
	fixedNoticeService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/notice/fixed"
	notFixedNoticeService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/notice/notFixed"
	suggestionService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/suggestion"
)

// GetOneBoardDetail 게시물 하나의 정보를 리턴
func (bc *BoardChaincode) GetOneBoardDetail(ctx contractapi.TransactionContextInterface, rawBoardGetOneRequest string) *response.ChaincodeResponse {
	var boardGetPostRequest request.BoardGetByKeyRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawBoardGetOneRequest,
	}

	err := json.Unmarshal([]byte(rawBoardGetOneRequest), &boardGetPostRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := boardService.GetOneBoardByKey(ctx, boardGetPostRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetBoardList Doctype에 따른 게시글 리스트를 리턴
func (bc *BoardChaincode) GetBoardList(ctx contractapi.TransactionContextInterface, rawBoardGetAllRequest string) *response.ChaincodeResponse {
	var boardGetAllRequest request.BoardGetAllRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawBoardGetAllRequest,
	}

	err := json.Unmarshal([]byte(rawBoardGetAllRequest), &boardGetAllRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := boardService.GetBoardListWithPaging(ctx, boardGetAllRequest)
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

// GetBoardListByDID 특정 DID 쓴 모든 건의사항들을 리턴
func (bc *BoardChaincode) GetBoardListByDID(ctx contractapi.TransactionContextInterface, rawBoardGetByDIDRequest string) *response.ChaincodeResponse {
	var boardGetByDIDRequest request.BoardGetByDIDRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawBoardGetByDIDRequest,
	}

	err := json.Unmarshal([]byte(rawBoardGetByDIDRequest), &boardGetByDIDRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := boardService.GetBoardListByDIDWithPaging(ctx, boardGetByDIDRequest)
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

// GetBoardListByWriter 특정 작성자가 쓴 모든 게시글들을 리턴
func (bc *BoardChaincode) GetBoardListByWriter(ctx contractapi.TransactionContextInterface, rawBoardGetByWriterRequest string) *response.ChaincodeResponse {
	var boardGetByWriterRequest request.BoardGetByWriterRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawBoardGetByWriterRequest,
	}

	err := json.Unmarshal([]byte(rawBoardGetByWriterRequest), &boardGetByWriterRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := boardService.GetBoardListByWriterWithPaging(ctx, boardGetByWriterRequest)
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

// GetOneNoticeDetail 공지 하나의 정보를 리턴
func (bc *BoardChaincode) GetOneNoticeDetail(ctx contractapi.TransactionContextInterface, rawNoticeGetOneRequest string) *response.ChaincodeResponse {
	var noticeGetPostRequest request.NoticeGetByKeyRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawNoticeGetOneRequest,
	}

	err := json.Unmarshal([]byte(rawNoticeGetOneRequest), &noticeGetPostRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}
	data, err := commonNoticeService.GetOneByKey(ctx, noticeGetPostRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetFixedNoticeList Doctype에 따른 게시글 리스트를 리턴
func (bc *BoardChaincode) GetFixedNoticeList(ctx contractapi.TransactionContextInterface) *response.ChaincodeResponse {
	chaincodeResponse := response.ChaincodeResponse{}
	data, err := fixedNoticeService.GetList(ctx)
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

// GetNotFixedNoticeList Doctype에 따른 게시글 리스트를 리턴
func (bc *BoardChaincode) GetNotFixedNoticeList(ctx contractapi.TransactionContextInterface, rawNoticeGetAllRequest string) *response.ChaincodeResponse {
	var noticeGetAllRequest request.NoticeGetAllRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawNoticeGetAllRequest,
	}

	err := json.Unmarshal([]byte(rawNoticeGetAllRequest), &noticeGetAllRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := notFixedNoticeService.GetListWithPaging(ctx, noticeGetAllRequest)
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

// GetOneSuggestionDetail 게시물 하나의 정보를 리턴
func (bc *BoardChaincode) GetOneSuggestionDetail(ctx contractapi.TransactionContextInterface, rawSuggestionGetOneRequest string) *response.ChaincodeResponse {
	var suggestionGetPostRequest request.SuggestionGetByKeyRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawSuggestionGetOneRequest,
	}

	err := json.Unmarshal([]byte(rawSuggestionGetOneRequest), &suggestionGetPostRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := suggestionService.GetOneSuggestionByKey(ctx, suggestionGetPostRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetSuggestionList Doctype에 따른 게시글 리스트를 리턴
func (bc *BoardChaincode) GetSuggestionList(ctx contractapi.TransactionContextInterface, rawSuggestionGetAllRequest string) *response.ChaincodeResponse {
	var suggestionGetAllRequest request.SuggestionGetAllRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawSuggestionGetAllRequest,
	}

	err := json.Unmarshal([]byte(rawSuggestionGetAllRequest), &suggestionGetAllRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := suggestionService.GetSuggestionListWithPaging(ctx, suggestionGetAllRequest)
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

// GetSuggestionListByDID 특정 DID 쓴 모든 건의사항들을 리턴
func (bc *BoardChaincode) GetSuggestionListByDID(ctx contractapi.TransactionContextInterface, rawSuggestionGetByDIDRequest string) *response.ChaincodeResponse {
	var suggestionGetByDIDRequest request.SuggestionGetByDIDRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawSuggestionGetByDIDRequest,
	}

	err := json.Unmarshal([]byte(rawSuggestionGetByDIDRequest), &suggestionGetByDIDRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := suggestionService.GetSuggestionListByDIDWithPaging(ctx, suggestionGetByDIDRequest)
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

// GetSuggestionListByWriter 특정 작성자가 쓴 모든 게시글들을 리턴
func (bc *BoardChaincode) GetSuggestionListByWriter(ctx contractapi.TransactionContextInterface, rawSuggestionGetByWriterRequest string) *response.ChaincodeResponse {
	var suggestionGetByWriterRequest request.SuggestionGetByWriterRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawSuggestionGetByWriterRequest,
	}

	err := json.Unmarshal([]byte(rawSuggestionGetByWriterRequest), &suggestionGetByWriterRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := suggestionService.GetSuggestionListByWriterWithPaging(ctx, suggestionGetByWriterRequest)
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

// GetOneCommunityDetail 게시물 하나의 정보를 리턴
func (bc *BoardChaincode) GetOneCommunityDetail(ctx contractapi.TransactionContextInterface, rawCommunityGetOneRequest string) *response.ChaincodeResponse {
	var communityGetPostRequest request.CommunityGetByKeyRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommunityGetOneRequest,
	}

	err := json.Unmarshal([]byte(rawCommunityGetOneRequest), &communityGetPostRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := communityService.GetOneCommunityByKey(ctx, communityGetPostRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetCommunityList Doctype에 따른 게시글 리스트를 리턴
func (bc *BoardChaincode) GetCommunityList(ctx contractapi.TransactionContextInterface, rawCommunityGetAllRequest string) *response.ChaincodeResponse {
	var communityGetAllRequest request.CommunityGetAllRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommunityGetAllRequest,
	}

	err := json.Unmarshal([]byte(rawCommunityGetAllRequest), &communityGetAllRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := communityService.GetCommunityListWithPaging(ctx, communityGetAllRequest)
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

// GetCommunityListByDID 특정 DID 쓴 모든 건의사항들을 리턴
func (bc *BoardChaincode) GetCommunityListByDID(ctx contractapi.TransactionContextInterface, rawCommunityGetByDIDRequest string) *response.ChaincodeResponse {
	var communityGetByDIDRequest request.CommunityGetByDIDRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommunityGetByDIDRequest,
	}

	err := json.Unmarshal([]byte(rawCommunityGetByDIDRequest), &communityGetByDIDRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := communityService.GetCommunityListByDIDWithPaging(ctx, communityGetByDIDRequest)
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

// GetCommunityListByWriter 특정 작성자가 쓴 모든 게시글들을 리턴
func (bc *BoardChaincode) GetCommunityListByWriter(ctx contractapi.TransactionContextInterface, rawCommunityGetByWriterRequest string) *response.ChaincodeResponse {
	var communityGetByWriterRequest request.CommunityGetByWriterRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawCommunityGetByWriterRequest,
	}

	err := json.Unmarshal([]byte(rawCommunityGetByWriterRequest), &communityGetByWriterRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := communityService.GetCommunityListByWriterWithPaging(ctx, communityGetByWriterRequest)
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

// GetOneInfoshareDetail 게시물 하나의 정보를 리턴
func (bc *BoardChaincode) GetOneInfoshareDetail(ctx contractapi.TransactionContextInterface, rawInfoshareGetOneRequest string) *response.ChaincodeResponse {
	var infoshareGetPostRequest request.InfoshareGetByKeyRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawInfoshareGetOneRequest,
	}

	err := json.Unmarshal([]byte(rawInfoshareGetOneRequest), &infoshareGetPostRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := infoshareService.GetOneInfoshareByKey(ctx, infoshareGetPostRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "OK"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetInfoshareList Doctype에 따른 게시글 리스트를 리턴
func (bc *BoardChaincode) GetInfoshareList(ctx contractapi.TransactionContextInterface, rawInfoshareGetAllRequest string) *response.ChaincodeResponse {
	var infoshareGetAllRequest request.InfoshareGetAllRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawInfoshareGetAllRequest,
	}

	err := json.Unmarshal([]byte(rawInfoshareGetAllRequest), &infoshareGetAllRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := infoshareService.GetInfoshareListWithPaging(ctx, infoshareGetAllRequest)
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

// GetInfoshareListByDID 특정 DID 쓴 모든 건의사항들을 리턴
func (bc *BoardChaincode) GetInfoshareListByDID(ctx contractapi.TransactionContextInterface, rawInfoshareGetByDIDRequest string) *response.ChaincodeResponse {
	var infoshareGetByDIDRequest request.InfoshareGetByDIDRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawInfoshareGetByDIDRequest,
	}

	err := json.Unmarshal([]byte(rawInfoshareGetByDIDRequest), &infoshareGetByDIDRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := infoshareService.GetInfoshareListByDIDWithPaging(ctx, infoshareGetByDIDRequest)
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

// GetInfoshareListByWriter 특정 작성자가 쓴 모든 게시글들을 리턴
func (bc *BoardChaincode) GetInfoshareListByWriter(ctx contractapi.TransactionContextInterface, rawInfoshareGetByWriterRequest string) *response.ChaincodeResponse {
	var infoshareGetByWriterRequest request.InfoshareGetByWriterRequest

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawInfoshareGetByWriterRequest,
	}

	err := json.Unmarshal([]byte(rawInfoshareGetByWriterRequest), &infoshareGetByWriterRequest)
	if err != nil {
		chaincodeResponse.Status = "유효하지 않은 DTO"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := infoshareService.GetInfoshareListByWriterWithPaging(ctx, infoshareGetByWriterRequest)
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
