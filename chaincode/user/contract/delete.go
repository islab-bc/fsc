package contract

import (
	"encoding/json"
	"net/http"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/request"
	response "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/response"
	userService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/service/user"
)

// DeleteUser ...
func (uc *UserChaincode) DeleteUser(ctx contractapi.TransactionContextInterface, rawUserDeleteRequest string) *response.ChaincodeResponse {
	// TODO: initRequest

	var userDeleteRequest request.UserDeleteRequest
	err := json.Unmarshal([]byte(rawUserDeleteRequest), &userDeleteRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserDeleteRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = userService.Delete(ctx, userDeleteRequest)

	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}
