package dto

import (
	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/common"
)

// NoticeCreateRequest ...
type NoticeCreateRequest struct {
	Company      string                `json:"company"`
	DID          string                `json:"did"`
	Writer       string                `json:"writer"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	Files        []common.FileMetadata `json:"files"`
	IsPopUp      uint16                `json:"is_pop_up"`
	PopUpContent string                `json:"pop_up_content"`
	EnrolledTime string                `json:"enrolled_time"`
}

// SuggestionCreateRequest ...
type SuggestionCreateRequest struct {
	Company      string                `json:"company"`
	DID          string                `json:"did"`
	Writer       string                `json:"writer"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	Files        []common.FileMetadata `json:"files"`
	Status       uint16                `json:"status"`
	EnrolledTime string                `json:"enrolled_time"`
}

// CommunityCreateRequest ...
type CommunityCreateRequest struct {
	Company      string                `json:"company"`
	DID          string                `json:"did"`
	Writer       string                `json:"writer"`
	Category     string                `json:"category"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	Files        []common.FileMetadata `json:"files"`
	EnrolledTime string                `json:"enrolled_time"`
}

// InfoshareCreateRequest ...
type InfoshareCreateRequest struct {
	Company      string                `json:"company"`
	DID          string                `json:"did"`
	Writer       string                `json:"writer"`
	Category     string                `json:"category"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	Files        []common.FileMetadata `json:"files"`
	EnrolledTime string                `json:"enrolled_time"`
}

// CommentCreateRequest is request DTO to add comment for board
type CommentCreateRequest struct {
	DID                string `json:"did"`
	BoardCompany       string `json:"board_company"`
	BoardEnrolledTime  string `json:"board_enrolled_time"`
	ParentCommenter    string `json:"parent_commenter"`
	ParentEnrolledTime string `json:"parent_enrolled_time"`
	Commenter          string `json:"commenter"`
	Content            string `json:"comment_content"`
	EnrolledTime       string `json:"comment_enrolled_time"`
}
