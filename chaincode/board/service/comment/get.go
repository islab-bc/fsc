package service

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/model"
)

// GetCommentsByBoardKey is function to get comments by writer with paging for board
func GetCommentsByBoardKey(ctx contractapi.TransactionContextInterface, BoardCompany, BoardEnrolledTime string) (string, error) {
	if BoardCompany == "" {
		logger.Error("게시글의 회사이름이 없음")
		return "", errors.New("게시글의 회사이름이 없음")
	}
	if BoardEnrolledTime == "" {
		logger.Error("게시글 등록시간 에러")
		return "", errors.New("게시글 등록시간 에러")
	}
	commentModel := model.NewComment(
		"comment",
	)
	commentModel.SetBoardCompany(BoardCompany)
	commentModel.SetBoardEnrolledTime(BoardEnrolledTime)
	boardKey := commentModel.GetBoardKey()
	boardKey = strings.Replace(boardKey, "(", "\\\\(", -1)
	boardKey = strings.Replace(boardKey, ")", "\\\\)", -1)
	queryString := `{"selector":{"_id": {"$regex":"` + boardKey + `"},"doctype":"comment"},"sort":[{"hierarchy_index":"desc"}],"use_index":["_design/indexCommentDoc","indexComment"]}`
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("iterator 정보를 가져오는데 실패함")
	}
	defer resultsIterator.Close()

	var comments []*model.Comment
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("다음 iterator 정보를 가져오는데 실패함")
		}
		var comment model.Comment
		err = json.Unmarshal(queryResponse.Value, &comment)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("JSON 디코딩에 실패함")
		}
		comments = append(comments, &comment)
	}
	for i, j := 0, len(comments)-1; i < j; i, j = i+1, j-1 {
		comments[i], comments[j] = comments[j], comments[i]
	}

	commentsBytes, err := json.Marshal(comments)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("JSON 인코딩에 실패함")
	}
	return string(commentsBytes), nil
}

// GetCommentsByBoardKeyWithPaging is function to get comments by writer with paging for board
func GetCommentsByBoardKeyWithPaging(ctx contractapi.TransactionContextInterface, cbr request.CommentGetByBoardKeyRequest) (string, error) {
	if cbr.BoardCompany == "" {
		logger.Error("게시글의 회사이름이 없음")
		return "", errors.New("게시글의 회사이름이 없음")
	}
	if cbr.BoardEnrolledTime == "" {
		logger.Error("게시글 등록시간 에러")
		return "", errors.New("게시글 등록시간 에러")
	}
	commentModel := model.NewComment(
		"comment",
	)
	commentModel.SetBoardCompany(cbr.BoardCompany)
	commentModel.SetBoardEnrolledTime(cbr.BoardEnrolledTime)
	boardKey := commentModel.GetBoardKey()
	if isBoardKeyAvailable(ctx, boardKey) == false {
		logger.Error("게시글을 찾을 수 없음")
		return "", errors.New("게시글을 찾을 수 없음")
	}
	queryString := `{"selector":{"_id": {"$regex":"` + boardKey + `"},"doctype":"comment"},"sort":[{"hierarchy_index":"desc"}],"use_index":["_design/indexCommentDoc","indexComment"]}`
	return getCommentListByQueryWithPaging(ctx, cbr.PageNum, cbr.PageSize, queryString)
}

// GetCommentListByDIDWithPaging is function to get comments by writer with paging for did
func GetCommentListByDIDWithPaging(ctx contractapi.TransactionContextInterface, ccr request.CommentGetByDIDRequest) (string, error) {
	did := ccr.DID
	if did == "" {
		logger.Error("DID가 없음")
		return "", errors.New("DID가 없음")
	}
	didList, err := InvokeGetDIDHistory(ctx, did)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("InvokeGetDIDHistory() 호출 실패")
	}
	queryString := `{"selector":{"$or":[`
	for i := 0; i < len(didList); i++ {
		if i == len(didList)-1 {
			queryString = queryString + `{"did":"` + didList[i] + `"}`
		} else {
			queryString = queryString + `{"did":"` + didList[i] + `"},`
		}
	}
	queryString = queryString + `],"doctype":"comment"},`
	queryString = queryString + `"use_index":["_design/indexCommentDoc","indexComment"]}`

	return getCommentListByQueryWithPaging(ctx, ccr.PageNum, ccr.PageSize, queryString)
}

// GetCommentsByCommenterWithPaging is function to get comments by writer with paging for commenter
func GetCommentsByCommenterWithPaging(ctx contractapi.TransactionContextInterface, cbr request.CommentGetByCommenterRequest) (string, error) {
	commenter := cbr.Commenter
	if commenter == "" {
		logger.Error("댓글 작성자가 없음")
		return "", errors.New("댓글 작성자가 없음")
	}
	queryString := `{"selector":{"commenter":"` + commenter + `","doctype":"comment"},"use_index":["_design/indexCommentDoc","indexComment"]}`
	return getCommentListByQueryWithPaging(ctx, cbr.PageNum, cbr.PageSize, queryString)
}

// getCommentListByQueryWithPaging is function to get board data with paging by query
func getCommentListByQueryWithPaging(ctx contractapi.TransactionContextInterface, num, size uint16, queryString string) (string, error) {
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

	var comments []*model.Comment
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("[페이징] nextIterator 에러")
		}

		var comment model.Comment
		err = json.Unmarshal(queryResponse.Value, &comment)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("[페이징] comment JSON 디코딩 에러")
		}
		comments = append(comments, &comment)
	}

	for i, j := 0, len(comments)-1; i < j; i, j = i+1, j-1 {
		comments[i], comments[j] = comments[j], comments[i]
	}

	count := meta.GetFetchedRecordsCount()
	pagingEntity := struct {
		Data  []*model.Comment `json:"comments"`
		Count int32            `json:"total_count"`
	}{Data: comments, Count: count}

	bufferedPagingEntity, err := json.Marshal(pagingEntity)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("[페이징] pageingEntity JSON 인코딩 에러")
	}

	return string(bufferedPagingEntity), nil
}

// getCommentByKey is function to get board data by key
func getCommentByKey(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	rawCommentData, err := ctx.GetStub().GetState(key)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("GetState 실패")
	}
	return string(rawCommentData), nil
}

// getBoardByKey is function to get board data by key
func getBoardByKey(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	rawBoardData, err := ctx.GetStub().GetState(key)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("GetState 실패")
	}
	return string(rawBoardData), nil
}

// GetCommentCount ...
func GetCommentCount(ctx contractapi.TransactionContextInterface, boardKey string) (uint16, error) {
	count := uint16(0)
	boardKey = strings.Replace(boardKey, "(", "\\\\(", -1)
	boardKey = strings.Replace(boardKey, ")", "\\\\)", -1)
	queryString := `{"selector":{"_id": {"$regex":"` + boardKey + `"},"doctype":"comment"},"use_index":["_design/indexCommentDoc","indexComment"]}`
	countIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		logger.Error(err.Error())
		return 0, errors.New("countIterator 에러")
	}
	defer countIterator.Close()

	for countIterator.HasNext() {
		countIterator.Next()
		count++
	}
	return count, nil
}

// InvokeGetDIDHistory ... invokeChaincode(user)
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

// CheckPermission ... invokeChaincode(user)
func CheckPermission(ctx contractapi.TransactionContextInterface, beforeDID string, didList []string) bool {
	for i := 0; i < len(didList); i++ {
		if didList[i] == beforeDID {
			return true
		}
	}
	return false
}
