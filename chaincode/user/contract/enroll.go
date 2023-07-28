package contract

import (
	"encoding/json"
	"net/http"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/request"
	response "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/response"
	userService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/service/user"
	admissionService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/service/admission"
)

// EnrollUser ...
func (uc *UserChaincode) EnrollUser(ctx contractapi.TransactionContextInterface, rawUserEnrollRequest string) *response.ChaincodeResponse {
	// TODO: initinitProcessing asserrion

	var userEnrollRequest request.UserEnrollRequest
	err := json.Unmarshal([]byte(rawUserEnrollRequest), &userEnrollRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserEnrollRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = userService.Enroll(ctx, userEnrollRequest)

	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}

// EnrollAdmission ...
func (uc *UserChaincode) EnrollAdmission(ctx contractapi.TransactionContextInterface, rawAdmissionEnrollRequest string) *response.ChaincodeResponse {
	// TODO: initinitProcessing asserrion

	var admissionEnrollRequest request.AdmissionEnrollRequest
	err := json.Unmarshal([]byte(rawAdmissionEnrollRequest), &admissionEnrollRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawAdmissionEnrollRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	err = admissionService.Enroll(ctx, admissionEnrollRequest)

	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
	}

	return &chaincodeResponse
}
