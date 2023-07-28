package model

import (
	"strings"

	"gitlab.smartm2m.co.kr/btp-testbed/chaincode/board/contract/common"
)

//Notice ...
type Notice struct {
	DocType      string                `json:"doctype"`
	Company      string                `json:"company"`
	DID          string                `json:"did"`
	Writer       string                `json:"writer"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	Files        []common.FileMetadata `json:"files"`
	IsPopUp      uint16                `json:"is_pop_up"`
	PopUpContent string                `json:"pop_up_content"`
	CommentCount uint16                `json:"comment_count"`
	EnrolledTime string                `json:"enrolled_time"`
	UpdatedTime  string                `json:"updated_time"`
}

//NewNotice ...
func NewNotice() *Notice {
	return &Notice{}
}

//GetDocType ...
func (n *Notice) GetDocType() string {
	return n.DocType
}

//GetCompany ...
func (n *Notice) GetCompany() string {
	return n.Company
}

//GetDID ...
func (n *Notice) GetDID() string {
	return n.DID
}

//GetWriter ...
func (n *Notice) GetWriter() string {
	return n.Writer
}

//GetTitle ...
func (n *Notice) GetTitle() string {
	return n.Title
}

//GetEnrolledTime ...
func (n *Notice) GetEnrolledTime() string {
	return n.EnrolledTime
}

//GetUpdatedTime ...
func (n *Notice) GetUpdatedTime() string {
	return n.UpdatedTime
}

//GetContent ...
func (n *Notice) GetContent() string {
	return n.Content
}

//GetFiles ...
func (n *Notice) GetFiles() []common.FileMetadata {
	return n.Files
}

//GetIsPopUp ...
func (n *Notice) GetIsPopUp() uint16 {
	return n.IsPopUp
}

//GetPopUpContent ...
func (n *Notice) GetPopUpContent() string {
	return n.PopUpContent
}

//GetCommentCount ...
func (n *Notice) GetCommentCount() uint16 {
	return n.CommentCount
}

//GetIFKey ...
func (n *Notice) GetIFKey() string {
	var sb strings.Builder
	sb.WriteString("IF_")
	sb.WriteString(n.Company)
	sb.WriteString("_")
	sb.WriteString(n.EnrolledTime)

	return sb.String()
}

//GetINFKey ...
func (n *Notice) GetINFKey() string {
	var sb strings.Builder
	sb.WriteString("INF_")
	sb.WriteString(n.Company)
	sb.WriteString("_")
	sb.WriteString(n.EnrolledTime)

	return sb.String()
}

//GetKeyForComment ...
func (n *Notice) GetKeyForComment() string {
	var sb strings.Builder
	sb.WriteString(n.Company)
	sb.WriteString("_")
	sb.WriteString(n.EnrolledTime)

	return sb.String()
}

//SetDocType ...
func (n *Notice) SetDocType(doctype string) error {
	n.DocType = doctype
	return nil
}

//SetCompany ...
func (n *Notice) SetCompany(company string) error {
	n.Company = company
	return nil
}

//SetDID ...
func (n *Notice) SetDID(did string) error {
	n.DID = did
	return nil
}

//SetWriter ...
func (n *Notice) SetWriter(writer string) error {
	n.Writer = writer
	return nil
}

//SetTitle ...
func (n *Notice) SetTitle(title string) error {
	n.Title = title
	return nil
}

//SetIsPopUp ...
func (n *Notice) SetIsPopUp(isPopUp uint16) error {
	n.IsPopUp = isPopUp
	return nil
}

//SetPopUpContent ...
func (n *Notice) SetPopUpContent(popUpContent string) error {
	n.PopUpContent = popUpContent
	return nil
}

//SetCommentCount ...
func (n *Notice) SetCommentCount(commentCount uint16) error {
	n.CommentCount = commentCount
	return nil
}

//SetContent ...
func (n *Notice) SetContent(content string) error {
	n.Content = content
	return nil
}

//SetFiles ...
func (n *Notice) SetFiles(files []common.FileMetadata) error {
	n.Files = files
	return nil
}

//SetEnrolledTime ...
func (n *Notice) SetEnrolledTime(enTime string) error {
	n.EnrolledTime = enTime
	return nil
}

//SetUpdatedTime ...
func (n *Notice) SetUpdatedTime(upTime string) error {
	n.UpdatedTime = upTime
	return nil
}
