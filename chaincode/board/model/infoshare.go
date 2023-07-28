package model

import (
	"strings"

	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/common"
)

//Infoshare ...
type Infoshare struct {
	DocType      string                `json:"doctype"`
	Company      string                `json:"company"`
	DID          string                `json:"did"`
	Writer       string                `json:"writer"`
	Category     string                `json:"category"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	Files        []common.FileMetadata `json:"files"`
	CommentCount uint16                `json:"comment_count"`
	EnrolledTime string                `json:"enrolled_time"`
	UpdatedTime  string                `json:"updated_time"`
}

//NewInfoshare ...
func NewInfoshare() *Infoshare {
	return &Infoshare{}
}

//GetDocType ...
func (i *Infoshare) GetDocType() string {
	return i.DocType
}

//GetCompany ...
func (i *Infoshare) GetCompany() string {
	return i.Company
}

//GetDID ...
func (i *Infoshare) GetDID() string {
	return i.DID
}

//GetWriter ...
func (i *Infoshare) GetWriter() string {
	return i.Writer
}

//GetCategory ...
func (i *Infoshare) GetCategory() string {
	return i.Category
}

//GetTitle ...
func (i *Infoshare) GetTitle() string {
	return i.Title
}

//GetEnrolledTime ...
func (i *Infoshare) GetEnrolledTime() string {
	return i.EnrolledTime
}

//GetUpdatedTime ...
func (i *Infoshare) GetUpdatedTime() string {
	return i.UpdatedTime
}

//GetContent ...
func (i *Infoshare) GetContent() string {
	return i.Content
}

//GetFiles ...
func (i *Infoshare) GetFiles() []common.FileMetadata {
	return i.Files
}

//GetCommentCount ...
func (i *Infoshare) GetCommentCount() uint16 {
	return i.CommentCount
}

//GetKey ...
func (i *Infoshare) GetKey() string {
	var sb strings.Builder
	sb.WriteString(i.Company)
	sb.WriteString("_")
	sb.WriteString(i.EnrolledTime)

	return sb.String()
}

//SetDocType ...
func (i *Infoshare) SetDocType(doctype string) error {
	i.DocType = doctype
	return nil
}

//SetCompany ...
func (i *Infoshare) SetCompany(company string) error {
	i.Company = company
	return nil
}

//SetDID ...
func (i *Infoshare) SetDID(did string) error {
	i.DID = did
	return nil
}

//SetWriter ...
func (i *Infoshare) SetWriter(writer string) error {
	i.Writer = writer
	return nil
}

//SetCategory ...
func (i *Infoshare) SetCategory(category string) error {
	i.Category = category
	return nil
}

//SetTitle ...
func (i *Infoshare) SetTitle(title string) error {
	i.Title = title
	return nil
}

//SetCommentCount ...
func (i *Infoshare) SetCommentCount(commentCount uint16) error {
	i.CommentCount = commentCount
	return nil
}

//SetContent ...
func (i *Infoshare) SetContent(content string) error {
	i.Content = content
	return nil
}

//SetFiles ...
func (i *Infoshare) SetFiles(files []common.FileMetadata) error {
	i.Files = files
	return nil
}

//SetEnrolledTime ...
func (i *Infoshare) SetEnrolledTime(enTime string) error {
	i.EnrolledTime = enTime
	return nil
}

//SetUpdatedTime ...
func (i *Infoshare) SetUpdatedTime(upTime string) error {
	i.UpdatedTime = upTime
	return nil
}
