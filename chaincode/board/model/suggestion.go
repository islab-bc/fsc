package model

import (
	"strings"

	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/common"
)

//Suggestion ...
type Suggestion struct {
	DocType      string                `json:"doctype"`
	Company      string                `json:"company"`
	DID          string                `json:"did"`
	Writer       string                `json:"writer"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	Files        []common.FileMetadata `json:"files"`
	Status       uint16                `json:"status"`
	CommentCount uint16                `json:"comment_count"`
	EnrolledTime string                `json:"enrolled_time"`
	UpdatedTime  string                `json:"updated_time"`
}

//NewSuggestion ...
func NewSuggestion() *Suggestion {
	return &Suggestion{}
}

//GetDocType ...
func (s *Suggestion) GetDocType() string {
	return s.DocType
}

//GetCompany ...
func (s *Suggestion) GetCompany() string {
	return s.Company
}

//GetDID ...
func (s *Suggestion) GetDID() string {
	return s.DID
}

//GetWriter ...
func (s *Suggestion) GetWriter() string {
	return s.Writer
}

//GetTitle ...
func (s *Suggestion) GetTitle() string {
	return s.Title
}

//GetEnrolledTime ...
func (s *Suggestion) GetEnrolledTime() string {
	return s.EnrolledTime
}

//GetUpdatedTime ...
func (s *Suggestion) GetUpdatedTime() string {
	return s.UpdatedTime
}

//GetContent ...
func (s *Suggestion) GetContent() string {
	return s.Content
}

//GetFiles ...
func (s *Suggestion) GetFiles() []common.FileMetadata {
	return s.Files
}

//GetCommentCount ...
func (s *Suggestion) GetCommentCount() uint16 {
	return s.CommentCount
}

//GetStatus ...
func (s *Suggestion) GetStatus() uint16 {
	return s.Status
}

//GetKey ...
func (s *Suggestion) GetKey() string {
	var sb strings.Builder
	sb.WriteString(s.Company)
	sb.WriteString("_")
	sb.WriteString(s.EnrolledTime)

	return sb.String()
}

//SetDocType ...
func (s *Suggestion) SetDocType(doctype string) error {
	s.DocType = doctype
	return nil
}

//SetDID ...
func (s *Suggestion) SetDID(did string) error {
	s.DID = did
	return nil
}

//SetCompany ...
func (s *Suggestion) SetCompany(company string) error {
	s.Company = company
	return nil
}

//SetWriter ...
func (s *Suggestion) SetWriter(writer string) error {
	s.Writer = writer
	return nil
}

//SetTitle ...
func (s *Suggestion) SetTitle(title string) error {
	s.Title = title
	return nil
}

//SetCommentCount ...
func (s *Suggestion) SetCommentCount(commentCount uint16) error {
	s.CommentCount = commentCount
	return nil
}

//SetContent ...
func (s *Suggestion) SetContent(content string) error {
	s.Content = content
	return nil
}

//SetFiles ...
func (s *Suggestion) SetFiles(files []common.FileMetadata) error {
	s.Files = files
	return nil
}

//SetStatus ...
func (s *Suggestion) SetStatus(status uint16) error {
	s.Status = status
	return nil
}

//SetEnrolledTime ...
func (s *Suggestion) SetEnrolledTime(enTime string) error {
	s.EnrolledTime = enTime
	return nil
}

//SetUpdatedTime ...
func (s *Suggestion) SetUpdatedTime(upTime string) error {
	s.UpdatedTime = upTime
	return nil
}
