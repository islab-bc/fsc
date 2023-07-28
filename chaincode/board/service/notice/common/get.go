package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/model"
	commentService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/comment"
	fixedNoticeService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/notice/fixed"
	notFixedNoticeService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/notice/notFixed"
)

// GetOneByKey 회사이름과 작성시간으로 공지를 하나 가져옴
func GetOneByKey(ctx contractapi.TransactionContextInterface, nkr request.NoticeGetByKeyRequest) (string, error) {
	if nkr.Company == "" {
		logger.Error("회사이름이 없음")
		return "", errors.New("회사이름이 없음")
	}
	if nkr.EnrolledTime == "" {
		logger.Error("등록시간 에러")
		return "", errors.New("등록시간 에러")
	}

	IFModel, _ := fixedNoticeService.GetOneByKey(ctx, nkr)
	INFModel, _ := notFixedNoticeService.GetOneByKey(ctx, nkr)

	noticeModelString := ""
	if len(IFModel) == 0 && len(INFModel) == 0 {
		logger.Error("게시글이 존재하지 않음")
		return "", errors.New("게시글이 존재하지 않음")
	}
	if len(IFModel) == 0 {
		noticeModelString = INFModel
		return noticeModelString, nil
	}
	if len(INFModel) == 0 {
		noticeModelString = IFModel
		return noticeModelString, nil
	}

	logger.Error("게시글이 fixed, not-fixed 모두 존재함")
	return "", errors.New("게시글이 fixed, not-fixed 모두 존재함")
}

// GetNoticeByKey is function to get board data by key
func GetNoticeByKey(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	rawBoardData, err := ctx.GetStub().GetState(key)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("GetState 실패")
	}
	return string(rawBoardData), nil
}

// GetListByQueryWithPaging is function to get board data with paging by query
func GetListByQueryWithPaging(ctx contractapi.TransactionContextInterface, num, size uint16, queryString string) (string, error) {
	bookmark := ""
	pageNum := int(num)
	if pageNum <= 0 {
		logger.Error("페이지 번호 에러")
		return "", errors.New("페이지 번호 에러")
	}
	pageSize := int32(size)
	if pageSize <= 0 {
		logger.Error("페이지 사이즈 에러")
		return "", errors.New("페이지 사이즈 에러")
	}

	for i := 0; i < pageNum-1; i++ {
		_, meta, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
		bookmark = meta.GetBookmark()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("[페이징] GetQueryResultWithPagination 실패")
		}
	}

	resultsIterator, meta, err := ctx.GetStub().GetQueryResultWithPagination(queryString, pageSize, bookmark)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("[페이징] resultsIterator 에러")
	}
	defer resultsIterator.Close()

	var notices []*model.Notice
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("[페이징] nextIterator 에러")
		}

		var notice model.Notice
		err = json.Unmarshal(queryResponse.Value, &notice)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("[페이징] board JSON 디코딩 에러")
		}
		commentCount, err := commentService.GetCommentCount(ctx, notice.GetKeyForComment())
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("[페이징] 게시글 별 댓글 개수 에러")
		}
		notice.SetCommentCount(commentCount)
		notices = append(notices, &notice)
	}

	count := meta.GetFetchedRecordsCount()
	pagingEntity := struct {
		Notices []*model.Notice `json:"notices"`
		Count   int32           `json:"total_count"`
	}{Notices: notices, Count: count}

	bufferedPagingEntity, err := json.Marshal(pagingEntity)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("[페이징] pageingEntity JSON 인코딩 에러")
	}

	return string(bufferedPagingEntity), nil
}

// InvokeGetDIDHistory ... invokeChaincode(notice)
func InvokeGetDIDHistory(ctx contractapi.TransactionContextInterface, currentDID string) ([]string, error) {
	channelName := "testbed"
	args := "{\"did\":\"" + currentDID + "\"}"
	params := []string{"GetDIDHistory", args}
	queryArgs := make([][]byte, len(params))
	for i, arg := range params {
		queryArgs[i] = []byte(arg)
	}
	didData := make(map[string]interface{})
	responseData := ctx.GetStub().InvokeChaincode("user", queryArgs, channelName)
	err := json.Unmarshal([]byte(responseData.Payload), &didData)
	if err != nil {
		logger.Error(err.Error())
		return []string{}, errors.New("JSON 디코딩에 실패함")
	}

	var didList []string
	err = json.Unmarshal([]byte(didData["data"].(string)), &didList)
	if err != nil {
		logger.Error(err.Error())
		return []string{}, errors.New("JSON 디코딩에 실패함")
	}
	return didList, nil
}

// CheckPermission ... invokeChaincode(notice)
func CheckPermission(ctx contractapi.TransactionContextInterface, beforeDID string, didList []string) bool {
	for i := 0; i < len(didList); i++ {
		if didList[i] == beforeDID {
			return true
		}
	}
	return false
}
