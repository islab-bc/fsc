package contract

import (
	"encoding/json"
	"net/http"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/request"
	response "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/response"
	userService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/service/user"
)

// UpdateDID ...
func (uc *UserChaincode) UpdateDID(ctx contractapi.TransactionContextInterface, rawUserUpdateRequest string) *response.ChaincodeResponse {
	// TODO: initinitProcessing asserrion

	var userUpdateRequest request.UserDIDUpdateRequest
	err := json.Unmarshal([]byte(rawUserUpdateRequest), &userUpdateRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserUpdateRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = userService.UpdateDID(ctx, userUpdateRequest)

	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// UpdateUser ...
func (uc *UserChaincode) UpdateUser(ctx contractapi.TransactionContextInterface, rawUserUpdateRequest string) *response.ChaincodeResponse {
	// TODO: initinitProcessing asserrion

	var userUpdateRequest request.UserDIDUpdateRequest
	err := json.Unmarshal([]byte(rawUserUpdateRequest), &userUpdateRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserUpdateRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = userService.UpdateUser(ctx, userUpdateRequest)

	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// UpdateUserState ...
func (uc *UserChaincode) UpdateUserState(ctx contractapi.TransactionContextInterface, rawUserUpdateRequest string) *response.ChaincodeResponse {
	// TODO: initinitProcessing asserrion

	var userUpdateRequest request.UserDIDUpdateRequest
	err := json.Unmarshal([]byte(rawUserUpdateRequest), &userUpdateRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserUpdateRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = userService.UpdateState(ctx, userUpdateRequest)

	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}
