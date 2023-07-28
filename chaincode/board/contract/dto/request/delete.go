package dto

// NoticeDeleteRequest is request DTO to delete post for board
type NoticeDeleteRequest struct {
	Company      string `json:"company"`
	EnrolledTime string `json:"enrolled_time"`
	DID			 string `json:"did"`
}

// SuggestionDeleteRequest is request DTO to delete post for board
type SuggestionDeleteRequest struct {
	DID			 string `json:"did"`
	Company      string `json:"company"`
	EnrolledTime string `json:"enrolled_time"`
}

// CommunityDeleteRequest is request DTO to delete post for board
type CommunityDeleteRequest struct {
	DID			 string `json:"did"`
	Company      string `json:"company"`
	EnrolledTime string `json:"enrolled_time"`
}

// InfoshareDeleteRequest is request DTO to delete post for board
type InfoshareDeleteRequest struct {
	DID			 string `json:"did"`
	Company      string `json:"company"`
	EnrolledTime string `json:"enrolled_time"`
}

// CommentDeleteRequest ...
type CommentDeleteRequest struct {
	DID			 	  string `json:"did"`
	BoardCompany      string `json:"board_company"`
	BoardEnrolledTime string `json:"board_enrolled_time"`
	Commenter         string `json:"commenter"`
	EnrolledTime      string `json:"enrolled_time"`
}
