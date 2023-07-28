package dto

// BoardGetByKeyRequest is request DTO to get post information for board
type BoardGetByKeyRequest struct {
	Company        string `json:"company"`
	EnrolledTime string `json:"enrolled_time"`
}

// BoardGetAllRequest ...
type BoardGetAllRequest struct {
	Query    string `json:"query"`
	Keyword  string `json:"keyword"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// BoardGetByDIDRequest ...
type BoardGetByDIDRequest struct {
	DID      string `json:"did"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// BoardGetByWriterRequest ...
type BoardGetByWriterRequest struct {
	Writer   string `json:"writer"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// NoticeGetByKeyRequest is request DTO to get post information for board
type NoticeGetByKeyRequest struct {
	Company      string `json:"company"`
	EnrolledTime string `json:"enrolled_time"`
}

// NoticeGetAllRequest ...
type NoticeGetAllRequest struct {
	Query    string `json:"query"`
	Keyword  string `json:"keyword"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// NoticeGetByDIDRequest ...
type NoticeGetByDIDRequest struct {
	DID      string `json:"did"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// NoticeGetByWriterRequest ...
type NoticeGetByWriterRequest struct {
	Writer   string `json:"writer"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// SuggestionGetByKeyRequest is request DTO to get post information for board
type SuggestionGetByKeyRequest struct {
	Company      string `json:"company"`
	EnrolledTime string `json:"enrolled_time"`
}

// SuggestionGetAllRequest ...
type SuggestionGetAllRequest struct {
	Query    string `json:"query"`
	Keyword  string `json:"keyword"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// SuggestionGetByDIDRequest ...
type SuggestionGetByDIDRequest struct {
	DID      string `json:"did"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// SuggestionGetByWriterRequest ...
type SuggestionGetByWriterRequest struct {
	Writer   string `json:"writer"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// CommunityGetByKeyRequest is request DTO to get post information for board
type CommunityGetByKeyRequest struct {
	Company      string `json:"company"`
	EnrolledTime string `json:"enrolled_time"`
}

// CommunityGetAllRequest ...
type CommunityGetAllRequest struct {
	Query    string `json:"query"`
	Keyword  string `json:"keyword"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// CommunityGetByDIDRequest ...
type CommunityGetByDIDRequest struct {
	DID      string `json:"did"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// CommunityGetByWriterRequest ...
type CommunityGetByWriterRequest struct {
	Writer   string `json:"writer"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// InfoshareGetByKeyRequest is request DTO to get post information for board
type InfoshareGetByKeyRequest struct {
	Company      string `json:"company"`
	EnrolledTime string `json:"enrolled_time"`
}

// InfoshareGetAllRequest ...
type InfoshareGetAllRequest struct {
	Query    string `json:"query"`
	Keyword  string `json:"keyword"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// InfoshareGetByDIDRequest ...
type InfoshareGetByDIDRequest struct {
	DID      string `json:"did"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// InfoshareGetByWriterRequest ...
type InfoshareGetByWriterRequest struct {
	Writer   string `json:"writer"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// CommentGetByBoardKeyRequest is request DTO to get list by writer for board
type CommentGetByBoardKeyRequest struct {
	BoardCompany      string `json:"board_company"`
	BoardEnrolledTime string `json:"board_enrolled_time"`
	PageNum           uint16 `json:"page_num"`
	PageSize          uint16 `json:"page_size"`
}

// CommentGetByDIDRequest is request DTO to get list by writer for board
type CommentGetByDIDRequest struct {
	DID      string `json:"did"`
	PageNum  uint16 `json:"page_num"`
	PageSize uint16 `json:"page_size"`
}

// CommentGetByCommenterRequest is request DTO to get list by writer for board
type CommentGetByCommenterRequest struct {
	Commenter string `json:"commenter"`
	PageNum   uint16 `json:"page_num"`
	PageSize  uint16 `json:"page_size"`
}
