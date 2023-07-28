package dto

import "gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/common"

// NoticeUpdateRequest is request DTO to update post for board
type NoticeUpdateRequest struct {
	Company      string                `json:"company"`
	DID          string                `json:"did"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	Files        []common.FileMetadata `json:"files"`
	IsFixed      uint16                `json:"is_fixed"`
	IsPopUp      uint16                `json:"is_pop_up"`
	PopUpContent string                `json:"pop_up_content"`
	EnrolledTime string                `json:"enrolled_time"`
	UpdatedTime  string                `json:"updated_time"`
}

// SuggestionUpdateRequest is request DTO to update post for board
type SuggestionUpdateRequest struct {
	Company      string                `json:"company"`
	DID          string                `json:"did"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	Files        []common.FileMetadata `json:"files"`
	Status       uint16                `json:"status"`
	EnrolledTime string                `json:"enrolled_time"`
	UpdatedTime  string                `json:"updated_time"`
}

// CommunityUpdateRequest is request DTO to update post for board
type CommunityUpdateRequest struct {
	Company      string                `json:"company"`
	DID          string                `json:"did"`
	Category     string                `json:"category"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	Files        []common.FileMetadata `json:"files"`
	EnrolledTime string                `json:"enrolled_time"`
	UpdatedTime  string                `json:"updated_time"`
}

// InfoshareUpdateRequest is request DTO to update post for board
type InfoshareUpdateRequest struct {
	Company      string                `json:"company"`
	DID          string                `json:"did"`
	Category     string                `json:"category"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	Files        []common.FileMetadata `json:"files"`
	EnrolledTime string                `json:"enrolled_time"`
	UpdatedTime  string                `json:"updated_time"`
}

// CommentUpdateRequest is request DTO to add comment for board
type CommentUpdateRequest struct {
	BoardCompany      string `json:"board_company"`
	DID               string `json:"did"`
	BoardEnrolledTime string `json:"board_enrolled_time"`
	Commenter         string `json:"commenter"`
	EnrolledTime      string `json:"enrolled_time"`
	Content           string `json:"comment_content"`
	UpdatedTime       string `json:"updated_time"`
}
