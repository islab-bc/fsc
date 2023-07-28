package service

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	request "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/dto/request"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/logger"
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/model"
	commentService "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/service/comment"
)

// GetOneBoardByKey 회사이름과 작성시간으로 글 하나 가져옴
func GetOneBoardByKey(ctx contractapi.TransactionContextInterface, bkr request.BoardGetByKeyRequest) (string, error) {
	if bkr.Company == "" {
		logger.Error("회사이름이 없음")
		return "", errors.New("회사이름이 없음")
	}
	if bkr.EnrolledTime == "" {
		logger.Error("등록시간 에러")
		return "", errors.New("등록시간 에러")
	}

	targetBoardModel := model.NewCommunity()
	targetBoardModel.SetCompany(bkr.Company)
	targetBoardModel.SetEnrolledTime(bkr.EnrolledTime)

	isCreated, err := getBoardByKey(ctx, targetBoardModel.GetKey())
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("GetState 실패")
	}
	if len(isCreated) == 0 {
		logger.Error("존재하지 않는 게시글")
		return "", errors.New("존재하지 않는 게시글")
	}
	err = json.Unmarshal([]byte(isCreated), &targetBoardModel)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("[페이징] board JSON 디코딩 에러")
	}
	commentsInfo, err := commentService.GetCommentsByBoardKey(ctx, bkr.Company, bkr.EnrolledTime)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("comment를 가져올 수 없음")
	}
	responseEntity := struct {
		Board        string `json:"board"`
		CommentsInfo string `json:"comments_info"`
	}{Board: isCreated, CommentsInfo: commentsInfo}

	bufferedResponseEntity, err := json.Marshal(responseEntity)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("responseEntity JSON 인코딩 에러")
	}
	return string(bufferedResponseEntity), nil
}

// GetBoardListByDIDWithPaging is function to get comments by writer with paging for did
func GetBoardListByDIDWithPaging(ctx contractapi.TransactionContextInterface, bdr request.BoardGetByDIDRequest) (string, error) {
	did := bdr.DID
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
			queryString = queryString + `{"did":"` + didList[i] + `"}]}`
		} else {
			queryString = queryString + `{"did":"` + didList[i] + `"},`
		}
	}
	queryString = queryString + `,"use_index":["_design/indexAllBoardDoc","indexAllBoard"]}`
	return getBoardListByQueryWithPaging(ctx, bdr.PageNum, bdr.PageSize, queryString)
}

// GetBoardListWithPaging 보드 리스트를 페이징해서 줌
func GetBoardListWithPaging(ctx contractapi.TransactionContextInterface, bar request.BoardGetAllRequest) (string, error) {
	query := bar.Query
	keyword := bar.Keyword
	if query == "" || keyword == "" {
		queryString := `{"selector":{"did":{"$gte":""}},"sort":[{"enrolled_time":"desc"}],"use_index":["_design/indexAllBoardDoc","indexAllBoard"]}`
		return getBoardListByQueryWithPaging(ctx, bar.PageNum, bar.PageSize, queryString)
	}
	queryString := `{"selector":{"` + query + `":"` + keyword + `"},"sort":[{"enrolled_time":"desc"}],"use_index":["_design/indexAllBoardDoc","indexAllBoard"]}`
	return getBoardListByQueryWithPaging(ctx, bar.PageNum, bar.PageSize, queryString)
}

// GetBoardListByWriterWithPaging 특정 작성자가 공지를 페이징해서 줌
func GetBoardListByWriterWithPaging(ctx contractapi.TransactionContextInterface, bwr request.BoardGetByWriterRequest) (string, error) {
	writer := bwr.Writer
	if writer == "" {
		logger.Error("작성자가 없음")
		return "", errors.New("작성자가 없음")
	}
	queryString := `{"selector":{"writer":"` + writer + `"},"sort":[{"enrolled_time":"desc"}],"use_index":["_design/indexAllBoardDoc","indexAllBoard"]}`
	return getBoardListByQueryWithPaging(ctx, bwr.PageNum, bwr.PageSize, queryString)
}

// getBoardListByQueryWithPaging is function to get board data with paging by query
func getBoardListByQueryWithPaging(ctx contractapi.TransactionContextInterface, num, size uint16, queryString string) (string, error) {
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

	var boards []*model.Community
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("[페이징] nextIterator 에러")
		}

		var board model.Community
		err = json.Unmarshal(queryResponse.Value, &board)
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("[페이징] board JSON 디코딩 에러")
		}
		commentCount, err := commentService.GetCommentCount(ctx, board.GetKey())
		if err != nil {
			logger.Error(err.Error())
			return "", errors.New("[페이징] 게시글 별 댓글 개수 에러")
		}
		board.SetCommentCount(commentCount)
		boards = append(boards, &board)
	}

	count := meta.GetFetchedRecordsCount()
	pagingEntity := struct {
		Boards []*model.Community `json:"boards"`
		Count  int32              `json:"total_count"`
	}{Boards: boards, Count: count}

	bufferedPagingEntity, err := json.Marshal(pagingEntity)
	if err != nil {
		logger.Error(err.Error())
		return "", errors.New("[페이징] pageingEntity JSON 인코딩 에러")
	}

	return string(bufferedPagingEntity), nil
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
