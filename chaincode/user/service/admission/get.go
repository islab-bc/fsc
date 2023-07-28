package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/model"
)

// getAdmissionByKey ...
func getAdmissionByKey(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	rawUserData, err := ctx.GetStub().GetPrivateData("collectionUsers", key)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("해당 사용자를 찾을 수 없음")
	}
	return string(rawUserData), nil
}

// GetList ...
func GetList(ctx contractapi.TransactionContextInterface) (string, error) {
	queryString := `{"selector":{"doctype":"admission"},"use_index":["_design/indexTimeDoc","indexTime"]}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionUsers", queryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()

	var admissions []*model.Admission
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("다음 iterator 정보를 가져오는데 실패함")
		}
		var admission model.Admission
		err = json.Unmarshal(queryResponse.Value, &admission)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}
		admissions = append(admissions, &admission)
	}

	admissionsBytes, err := json.Marshal(admissions)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}

	return string(admissionsBytes), nil
}

// GetListByDID ...
func GetListByDID(ctx contractapi.TransactionContextInterface, aur request.AdmissionGetByUserRequest) (string, error) {
	if aur.DID == "" {
		logger.Error("유효하지 않은 DID")
    	return "", errors.New("유효하지 않은 DID")
	}
	queryString := `{"selector":{"doctype":"admission","did":"`+ aur.DID +`"},"use_index":["_design/indexTimeDoc","indexTime"]}`
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult("collectionUsers", queryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()

	var admissions []*model.Admission
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("다음 iterator 정보를 가져오는데 실패함")
		}
		var admission model.Admission
		err = json.Unmarshal(queryResponse.Value, &admission)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}
		admissions = append(admissions, &admission)
	}

	admissionsBytes, err := json.Marshal(admissions)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}

	return string(admissionsBytes), nil
}