package contract

import (
	"encoding/json"
	"net/http"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	response "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/response"
)

//BoardChaincode ...
type BoardChaincode struct {
	contractapi.Contract
}

//InitProcessing ...
func InitProcessing(rawBoardRequest string) (interface{}, response.ChaincodeResponse) {
	var boardRequest interface{}
	boilerPlateResponse := response.ChaincodeResponse{
		ChaincodeResult: rawBoardRequest,
	}

	err := json.Unmarshal([]byte(rawBoardRequest), &boardRequest)

	if err != nil {
		boilerPlateResponse.Status = "초기화 작업 실패"
		boilerPlateResponse.StatusCode = http.StatusInternalServerError
	}

	return boardRequest, boilerPlateResponse
}
