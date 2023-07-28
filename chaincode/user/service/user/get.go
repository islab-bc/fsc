package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/model"
)

// getUserByKey ...
func getUserByKey(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	rawUserData, err := ctx.GetStub().GetPrivateData("collectionUsers", key)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("해당 사용자를 찾을 수 없음")
	}
	return string(rawUserData), nil
}

// toChaincodeArgs converts string args to []byte args
func toChaincodeArgs(args ...string) [][]byte {
	bargs := make([][]byte, len(args))
	for i, arg := range args {
		bargs[i] = []byte(arg)
	}
	return bargs
}

// getLocationByCompanyName ...
func getLocationByCompanyName(ctx contractapi.TransactionContextInterface, channelName, companyName string) (string, error) {
	args := toChaincodeArgs("GetLocationByCompanyName", `{"company_name":"`+companyName+`"}`)

	response := ctx.GetStub().InvokeChaincode("company", args, channelName)

	company := response.Payload

	var locationObject struct {
		Data string `json:data`
	}

	err := json.Unmarshal(company, &locationObject)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return locationObject.Data, nil
}

// GetList ...
func GetList(ctx contractapi.TransactionContextInterface) (string, error) {
	queryString := `{"selector":{"doctype":"user"},"use_index":["_design/indexTimeDoc","indexTime"]}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionUsers", queryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()

	var users []*struct {
		Data model.User `json:"user"`
		Key  string     `json:"key"`
	}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("다음 iterator 정보를 가져오는데 실패함")
		}
		var user model.User
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}

		userData := struct {
			Data model.User `json:"user"`
			Key  string     `json:"key"`
		}{Data: user, Key: queryResponse.GetKey()}
		users = append(users, &userData)
	}

	usersBytes, err := json.Marshal(users)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}

	return string(usersBytes), nil
}

// GetApprovedList ...
func GetApprovedList(ctx contractapi.TransactionContextInterface) (string, error) {
	queryString := `{"selector":{"doctype":"user", "state":1},"use_index":["_design/indexTimeDoc","indexTime"]}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionUsers", queryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()

	var users []*struct {
		Data model.User `json:"user"`
		Key  string     `json:"key"`
	}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("다음 iterator 정보를 가져오는데 실패함")
		}
		var user model.User
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}
		userData := struct {
			Data model.User `json:"user"`
			Key  string     `json:"key"`
		}{Data: user, Key: queryResponse.GetKey()}
		users = append(users, &userData)
	}

	usersBytes, err := json.Marshal(users)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}

	return string(usersBytes), nil
}

// GetPendingList ...
func GetPendingList(ctx contractapi.TransactionContextInterface) (string, error) {
	queryString := `{"selector":{"doctype":"user", "state":0},"use_index":["_design/indexTimeDoc","indexTime"]}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionUsers", queryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()

	var users []*struct {
		Data model.User `json:"user"`
		Key  string     `json:"key"`
	}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("다음 iterator 정보를 가져오는데 실패함")
		}
		var user model.User
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}
		userData := struct {
			Data model.User `json:"user"`
			Key  string     `json:"key"`
		}{Data: user, Key: queryResponse.GetKey()}
		users = append(users, &userData)
	}

	usersBytes, err := json.Marshal(users)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}

	return string(usersBytes), nil
}

// GetRejectedList ...
func GetRejectedList(ctx contractapi.TransactionContextInterface) (string, error) {
	queryString := `{"selector":{"doctype":"user", "state":2},"use_index":["_design/indexTimeDoc","indexTime"]}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionUsers", queryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()

	var users []*struct {
		Data model.User `json:"user"`
		Key  string     `json:"key"`
	}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("다음 iterator 정보를 가져오는데 실패함")
		}
		var user model.User
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}
		userData := struct {
			Data model.User `json:"user"`
			Key  string     `json:"key"`
		}{Data: user, Key: queryResponse.GetKey()}
		users = append(users, &userData)
	}

	usersBytes, err := json.Marshal(users)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}

	return string(usersBytes), nil
}

// GetExpiredList ...
func GetExpiredList(ctx contractapi.TransactionContextInterface) (string, error) {
	queryString := `{"selector":{"doctype":"user", "state":3},"use_index":["_design/indexTimeDoc","indexTime"]}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionUsers", queryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()

	var users []*struct {
		Data model.User `json:"user"`
		Key  string     `json:"key"`
	}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("다음 iterator 정보를 가져오는데 실패함")
		}
		var user model.User
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}
		userData := struct {
			Data model.User `json:"user"`
			Key  string     `json:"key"`
		}{Data: user, Key: queryResponse.GetKey()}
		users = append(users, &userData)
	}

	usersBytes, err := json.Marshal(users)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}

	return string(usersBytes), nil
}

// GetAdminList ...
func GetAdminList(ctx contractapi.TransactionContextInterface) (string, error) {
	queryString := `{"selector":{"doctype":"user", "state":4},"use_index":["_design/indexTimeDoc","indexTime"]}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionUsers", queryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()

	var users []*struct {
		Data model.User `json:"user"`
		Key  string     `json:"key"`
	}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("다음 iterator 정보를 가져오는데 실패함")
		}
		var user model.User
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}
		userData := struct {
			Data model.User `json:"user"`
			Key  string     `json:"key"`
		}{Data: user, Key: queryResponse.GetKey()}
		users = append(users, &userData)
	}

	usersBytes, err := json.Marshal(users)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}

	return string(usersBytes), nil
}

// GetByQuery ...
func GetByQuery(ctx contractapi.TransactionContextInterface, uqr request.UserGetByQueryRequest) (string, error) {
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionUsers", uqr.QueryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()

	var users []*model.User
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("다음 iterator 정보를 가져오는데 실패함")
		}

		var user model.User
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}
		users = append(users, &user)
	}
	usersBytes, err := json.Marshal(users)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}

	return string(usersBytes), nil
}

// GetUserInfo ...
func GetUserInfo(ctx contractapi.TransactionContextInterface, uur request.UserGetByUserRequest) (string, error) {
	if uur.DID == "" {
		logger.Error("유효하지 않은 DID")
		return "", errors.New("유효하지 않은 DID")
	}
	queryString := `{"selector":{"doctype":"user"}, "use_index":["_design/indexCiDoc","indexCi"]}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionUsers", queryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()

	var result model.User
	channelName := "testbed"

	for resultsIterator.HasNext() {
		var user model.User
		queryResponse, err := resultsIterator.Next()
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}
		if queryResponse.GetKey() == uur.DID {
			err = json.Unmarshal(queryResponse.Value, &result)

			location, err := getLocationByCompanyName(ctx, channelName, result.CompanyName)
			if err != nil {
				logger.Error(err.Error())
				return "", err
			}

			result.SetLocation(location)

			resultData := struct {
				Data model.User `json:"user"`
				Key  string     `json:"key"`
			}{Data: result, Key: queryResponse.GetKey()}

			usersBytes, err := json.Marshal(resultData)
			if err != nil {
				logger.Error(err.Error())
				return "", errors.New("JSON 인코딩에 실패함")
			}
			return string(usersBytes), nil
		}
	}

	return "", errors.New("등록되지 않은 사용자")
}

// GetDeviceTokenByUserRequest ...
func GetDeviceTokenByUserRequest(ctx contractapi.TransactionContextInterface, uur request.UserGetByUserRequest) (string, error) {
	rawUserData, err := ctx.GetStub().GetPrivateData("collectionUsers", uur.DID)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("해당 사용자를 찾을 수 없음")
	}
	if len(rawUserData) == 0 {
		logger.Error("해당 사용자를 찾을 수 없음")
		return "", errors.New("해당 사용자를 찾을 수 없음")
	}
	var user model.User
	err = json.Unmarshal([]byte(string(rawUserData)), &user)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 디코딩에 실패함")
	}
	deviceToken := user.DeviceToken
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}
	return deviceToken, nil
}

// GetCIByUserRequest ...
func GetCIByUserRequest(ctx contractapi.TransactionContextInterface, uur request.UserGetByUserRequest) (string, error) {
	rawUserData, err := ctx.GetStub().GetPrivateData("collectionUsers", uur.DID)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("해당 사용자를 찾을 수 없음")
	}
	if len(rawUserData) == 0 {
		logger.Error("해당 사용자를 찾을 수 없음")
		return "", errors.New("해당 사용자를 찾을 수 없음")
	}
	var user model.User
	err = json.Unmarshal([]byte(string(rawUserData)), &user)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 디코딩에 실패함")
	}
	ci := user.CI
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}
	return ci, nil
}

// GetUserByCI ...
func GetUserByCI(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	if key == "" {
		logger.Error("유효하지 않은 CI")
		return "", errors.New("유효하지 않은 CI")
	}
	queryString := `{"selector":{"doctype":"user","ci":"` + key + `"},"use_index":["_design/indexCiDoc","indexCi"]}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionUsers", queryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()
	var result model.User

	for resultsIterator.HasNext() {
		var user model.User
		queryResponse, err := resultsIterator.Next()
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}
		err = json.Unmarshal(queryResponse.Value, &result)
		resultData := struct {
			Data model.User `json:"user"`
			Key  string     `json:"key"`
		}{Data: result, Key: queryResponse.GetKey()}
		usersBytes, err := json.Marshal(resultData)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 인코딩에 실패함")
		}
		return string(usersBytes), nil
	}

	return "", errors.New("등록되지 않은 사용자")
}

// GetProfileImageByUserRequest ...
func GetProfileImageByUserRequest(ctx contractapi.TransactionContextInterface, uur request.UserGetByUserRequest) (string, error) {
	rawUserData, err := ctx.GetStub().GetPrivateData("collectionUsers", uur.DID)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("해당 사용자를 찾을 수 없음")
	}
	if len(rawUserData) == 0 {
		logger.Error("해당 사용자를 찾을 수 없음")
		return "", errors.New("해당 사용자를 찾을 수 없음")
	}
	var user model.User
	err = json.Unmarshal([]byte(string(rawUserData)), &user)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 디코딩에 실패함")
	}
	imagePath := user.ImagePath
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}
	return imagePath, nil
}

// GetUserByPhoneNumber ...
func GetUserByPhoneNumber(ctx contractapi.TransactionContextInterface, upr request.UserGetByPhoneRequest) (string, error) {
	if upr.PhoneNumber == "" {
		logger.Error("휴대폰 번호가 입력되지 않음")
		return "", errors.New("휴대폰 번호가 입력되지 않음")
	}
	queryString := `{"selector":{"doctype":"user","phone_number":"` + upr.PhoneNumber + `"},"use_index":["_design/indexPhoneDoc","indexPhone"]}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionUsers", queryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()
	var result model.User

	for resultsIterator.HasNext() {
		var user model.User
		queryResponse, err := resultsIterator.Next()
		err = json.Unmarshal(queryResponse.Value, &user)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}
		err = json.Unmarshal(queryResponse.Value, &result)
		resultData := struct {
			Data model.User `json:"user"`
			Key  string     `json:"key"`
		}{Data: result, Key: queryResponse.GetKey()}
		usersBytes, err := json.Marshal(resultData)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 인코딩에 실패함")
		}
		return string(usersBytes), nil
	}

	return "", errors.New("등록되지 않은 사용자")
}

// GetDIDHistory ...
func GetDIDHistory(ctx contractapi.TransactionContextInterface, uur request.UserGetByUserRequest) (string, error) {
	var didList []string
	didList = append(didList, uur.DID)
	rawDIDData, err := ctx.GetStub().GetState(uur.DID)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("GetState 실패")
	}
	for string(rawDIDData) != "" {
		didList = append(didList, string(rawDIDData))
		rawDIDData, err = ctx.GetStub().GetState(string(rawDIDData))
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("GetState 실패")
		}
	}
	didBytes, err := json.Marshal(didList)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}
	return string(didBytes), nil
}
