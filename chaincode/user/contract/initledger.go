package contract

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/user/model"
)

// InitLedger ...
func (uc *UserChaincode) InitLedger(ctx contractapi.TransactionContextInterface) error {
	isInitBytes, err := ctx.GetStub().GetPrivateData("collectionUsers", "isInit")
	if err != nil {
		return fmt.Errorf("GetPrivateData() 에러 발생")
	} else if isInitBytes == nil {
		var initUser = model.User{
			"user", "admin", "admin", "name1", "name1", "name1", "name1", "name1", "name1", "name1", "name1", 4, "admin", "admin",
		}

		initUserAsBytes, err := json.Marshal(initUser)
		ctx.GetStub().PutPrivateData("collectionUsers", "admin", initUserAsBytes)
		if err != nil {
			return fmt.Errorf("저장에 실패함 %v", err)
		}
		if err != nil {
			return fmt.Errorf("JSON 인코딩에 실패함 %v", err)
		}

		return nil
	} else {
		return fmt.Errorf("이미 초기화되었음")
	}
}
