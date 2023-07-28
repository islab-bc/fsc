package model

import (
	"strings"

	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/common"
)

//Community ...
type Community struct {
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

//NewCommunity ...
func NewCommunity() *Community {
	return &Community{}
}

//GetDocType ...
func (c *Community) GetDocType() string {
	return c.DocType
}

//GetDID ...
func (c *Community) GetDID() string {
	return c.DID
}

//GetCompany ...
func (c *Community) GetCompany() string {
	return c.Company
}

//GetWriter ...
func (c *Community) GetWriter() string {
	return c.Writer
}

//GetCategory ...
func (c *Community) GetCategory() string {
	return c.Category
}

//GetTitle ...
func (c *Community) GetTitle() string {
	return c.Title
}

//GetEnrolledTime ...
func (c *Community) GetEnrolledTime() string {
	return c.EnrolledTime
}

//GetUpdatedTime ...
func (c *Community) GetUpdatedTime() string {
	return c.UpdatedTime
}

//GetContent ...
func (c *Community) GetContent() string {
	return c.Content
}

//GetFiles ...
func (c *Community) GetFiles() []common.FileMetadata {
	return c.Files
}

//GetCommentCount ...
func (c *Community) GetCommentCount() uint16 {
	return c.CommentCount
}

//GetKey ...
func (c *Community) GetKey() string {
	var sb strings.Builder
	sb.WriteString(c.Company)
	sb.WriteString("_")
	sb.WriteString(c.EnrolledTime)

	return sb.String()
}

//SetDocType ...
func (c *Community) SetDocType(doctype string) error {
	c.DocType = doctype
	return nil
}

//SetCompany ...
func (c *Community) SetCompany(company string) error {
	c.Company = company
	return nil
}

//SetDID ...
func (c *Community) SetDID(did string) error {
	c.DID = did
	return nil
}

//SetWriter ...
func (c *Community) SetWriter(writer string) error {
	c.Writer = writer
	return nil
}

//SetCategory ...
func (c *Community) SetCategory(category string) error {
	c.Category = category
	return nil
}

//SetTitle ...
func (c *Community) SetTitle(title string) error {
	c.Title = title
	return nil
}

//SetCommentCount ...
func (c *Community) SetCommentCount(commentCount uint16) error {
	c.CommentCount = commentCount
	return nil
}

//SetContent ...
func (c *Community) SetContent(content string) error {
	c.Content = content
	return nil
}

//SetFiles ...
func (c *Community) SetFiles(files []common.FileMetadata) error {
	c.Files = files
	return nil
}

//SetEnrolledTime ...
func (c *Community) SetEnrolledTime(enTime string) error {
	c.EnrolledTime = enTime
	return nil
}

//SetUpdatedTime ...
func (c *Community) SetUpdatedTime(upTime string) error {
	c.UpdatedTime = upTime
	return nil
}
