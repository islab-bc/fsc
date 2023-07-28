package contract

import (
	"encoding/json"
	"net/http"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	response "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/response"
)

// UserChaincode ...
type UserChaincode struct {
	contractapi.Contract
}

// InitProcessing ...
func InitProcessing(rawUserRequest string) (interface{}, response.ChaincodeResponse) {
	var userRequest interface{}
	boilerPlateResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserRequest,
	}

	err := json.Unmarshal([]byte(rawUserRequest), &userRequest)

	if err != nil {
		boilerPlateResponse.Status = "초기화 작업에 실패함"
		boilerPlateResponse.StatusCode = http.StatusInternalServerError
	}

	return userRequest, boilerPlateResponse
}
