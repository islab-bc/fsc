package contract

import (
	"encoding/json"
	"net/http"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/request"
	response "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/response"
	admissionService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/service/admission"
	userService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/service/user"
)

// GetUserList ...
func (uc *UserChaincode) GetUserList(ctx contractapi.TransactionContextInterface) *response.ChaincodeResponse {
	chaincodeResponse := response.ChaincodeResponse{}

	data, err := userService.GetList(ctx)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetApprovedUserList ...
func (uc *UserChaincode) GetApprovedUserList(ctx contractapi.TransactionContextInterface) *response.ChaincodeResponse {
	chaincodeResponse := response.ChaincodeResponse{}

	data, err := userService.GetApprovedList(ctx)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetUserList ...
func (uc *UserChaincode) GetPendingUserList(ctx contractapi.TransactionContextInterface) *response.ChaincodeResponse {
	chaincodeResponse := response.ChaincodeResponse{}

	data, err := userService.GetPendingList(ctx)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetRejectedUserList ...
func (uc *UserChaincode) GetRejectedUserList(ctx contractapi.TransactionContextInterface) *response.ChaincodeResponse {
	chaincodeResponse := response.ChaincodeResponse{}

	data, err := userService.GetRejectedList(ctx)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetExpiredUserList ...
func (uc *UserChaincode) GetExpiredUserList(ctx contractapi.TransactionContextInterface) *response.ChaincodeResponse {
	chaincodeResponse := response.ChaincodeResponse{}

	data, err := userService.GetExpiredList(ctx)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetAdminUserList ...
func (uc *UserChaincode) GetAdminUserList(ctx contractapi.TransactionContextInterface) *response.ChaincodeResponse {
	chaincodeResponse := response.ChaincodeResponse{}

	data, err := userService.GetAdminList(ctx)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetUserListByQuery ...
func (uc *UserChaincode) GetUserListByQuery(ctx contractapi.TransactionContextInterface, rawUserGetByQueryRequest string) *response.ChaincodeResponse {
	// TODO: initRequest

	var userGetByQueryRequest request.UserGetByQueryRequest
	err := json.Unmarshal([]byte(rawUserGetByQueryRequest), &userGetByQueryRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserGetByQueryRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := userService.GetByQuery(ctx, userGetByQueryRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetUserInfo ...
func (uc *UserChaincode) GetUserProfile(ctx contractapi.TransactionContextInterface, rawUserGetByUserRequest string) *response.ChaincodeResponse {
	// TODO: initRequest

	var userGetByUserRequest request.UserGetByUserRequest
	err := json.Unmarshal([]byte(rawUserGetByUserRequest), &userGetByUserRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserGetByUserRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, getUserInfoError := userService.GetUserInfo(ctx, userGetByUserRequest)
	if getUserInfoError != nil {
		chaincodeResponse.Status = getUserInfoError.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}
	return &chaincodeResponse
}

// GetDeviceToken ...
func (uc *UserChaincode) GetUserDeviceToken(ctx contractapi.TransactionContextInterface, rawUserGetByUserRequest string) *response.ChaincodeResponse {
	// TODO: initRequest

	var userGetByUserRequest request.UserGetByUserRequest
	err := json.Unmarshal([]byte(rawUserGetByUserRequest), &userGetByUserRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserGetByUserRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := userService.GetDeviceTokenByUserRequest(ctx, userGetByUserRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetUserProfileImage ...
func (uc *UserChaincode) GetUserProfileImage(ctx contractapi.TransactionContextInterface, rawUserGetByUserRequest string) *response.ChaincodeResponse {
	// TODO: initRequest

	var userGetByUserRequest request.UserGetByUserRequest
	err := json.Unmarshal([]byte(rawUserGetByUserRequest), &userGetByUserRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserGetByUserRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := userService.GetProfileImageByUserRequest(ctx, userGetByUserRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetUserCI ...
func (uc *UserChaincode) GetUserCI(ctx contractapi.TransactionContextInterface, rawUserGetByUserRequest string) *response.ChaincodeResponse {
	// TODO: initRequest

	var userGetByUserRequest request.UserGetByUserRequest
	err := json.Unmarshal([]byte(rawUserGetByUserRequest), &userGetByUserRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserGetByUserRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := userService.GetCIByUserRequest(ctx, userGetByUserRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetUserByCI ...
func (uc *UserChaincode) GetUserByCI(ctx contractapi.TransactionContextInterface, rawUserGetByUpdateRequest string) *response.ChaincodeResponse {
	// TODO: initRequest

	var userGetByUpdateRequest request.UserGetByUpdateRequest
	err := json.Unmarshal([]byte(rawUserGetByUpdateRequest), &userGetByUpdateRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserGetByUpdateRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := userService.GetUserByCI(ctx, userGetByUpdateRequest.CI)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetAdmissionList ...
func (uc *UserChaincode) GetAdmissionList(ctx contractapi.TransactionContextInterface) *response.ChaincodeResponse {
	chaincodeResponse := response.ChaincodeResponse{}

	data, err := admissionService.GetList(ctx)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetAdmissionListByDID ...
func (uc *UserChaincode) GetAdmissionListByDID(ctx contractapi.TransactionContextInterface, rawAdmissionGetByUserRequest string) *response.ChaincodeResponse {
	// TODO: initRequest

	var admissionGetByUserRequest request.AdmissionGetByUserRequest
	err := json.Unmarshal([]byte(rawAdmissionGetByUserRequest), &admissionGetByUserRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawAdmissionGetByUserRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := admissionService.GetListByDID(ctx, admissionGetByUserRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetUserByPhoneNumber ...
func (uc *UserChaincode) GetUserByPhoneNumber(ctx contractapi.TransactionContextInterface, rawUserGetByPhoneRequest string) *response.ChaincodeResponse {
	// TODO: initRequest

	var userGetByPhoneRequest request.UserGetByPhoneRequest
	err := json.Unmarshal([]byte(rawUserGetByPhoneRequest), &userGetByPhoneRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserGetByPhoneRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := userService.GetUserByPhoneNumber(ctx, userGetByPhoneRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}

// GetDIDHistory ...
func (uc *UserChaincode) GetDIDHistory(ctx contractapi.TransactionContextInterface, rawUserGetByUserRequest string) *response.ChaincodeResponse {
	// TODO: initRequest

	var userGetByUserRequest request.UserGetByUserRequest
	err := json.Unmarshal([]byte(rawUserGetByUserRequest), &userGetByUserRequest)

	chaincodeResponse := response.ChaincodeResponse{
		ChaincodeResult: rawUserGetByUserRequest,
	}

	if err != nil {
		chaincodeResponse.Status = "DTO가 유효하지 않음"
		chaincodeResponse.StatusCode = http.StatusInternalServerError
		return &chaincodeResponse
	}

	data, err := userService.GetDIDHistory(ctx, userGetByUserRequest)
	if err != nil {
		chaincodeResponse.Status = err.Error()
		chaincodeResponse.StatusCode = http.StatusInternalServerError
	} else {
		chaincodeResponse.Status = "성공"
		chaincodeResponse.StatusCode = http.StatusOK
		chaincodeResponse.ChaincodeResult = data
	}

	return &chaincodeResponse
}
