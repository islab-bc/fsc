package model

import "strings"

//Comment ...
type Comment struct {
	DocType            string `json:"doctype"`
	DID                string `json:"did"`
	BoardCompany       string `json:"board_company"`
	BoardEnrolledTime  string `json:"board_enrolled_time"`
	ParentCommenter    string `json:"parent_commenter"`
	ParentEnrolledTime string `json:"parent_enrolled_time"`
	Commenter          string `json:"commenter"`
	Content            string `json:"comment_content"`
	HierarchyIndex     string `json:"hierarchy_index"`
	EnrolledTime       string `json:"enrolled_time"`
	UpdatedTime        string `json:"updated_time"`
}

//NewComment ...
func NewComment(doctype string) *Comment {
	return &Comment{
		DocType: doctype,
	}
}

//GetDID ...
func (c *Comment) GetDID() string {
	return c.DID
}

//GetBoardCompany ...
func (c *Comment) GetBoardCompany() string {
	return c.BoardCompany
}

//GetBoardEnrolledTime ...
func (c *Comment) GetBoardEnrolledTime() string {
	return c.BoardEnrolledTime
}

//GetParentCommenter ...
func (c *Comment) GetParentCommenter() string {
	return c.ParentCommenter
}

//GetParentEnrolledTime ...
func (c *Comment) GetParentEnrolledTime() string {
	return c.ParentCommenter
}

//GetCommenter ...
func (c *Comment) GetCommenter() string {
	return c.Commenter
}

//GetContent ...
func (c *Comment) GetContent() string {
	return c.Content
}

//GetEnrolledTime ...
func (c *Comment) GetEnrolledTime() string {
	return c.EnrolledTime
}

//GetUpdatedTime ...
func (c *Comment) GetUpdatedTime() string {
	return c.UpdatedTime
}

//GetHierarchyIndex ...
func (c *Comment) GetHierarchyIndex() string {
	return c.HierarchyIndex
}

//GetKey ...
func (c *Comment) GetKey() string {
	var sb strings.Builder
	sb.WriteString(c.BoardCompany)
	sb.WriteString("_")
	sb.WriteString(c.BoardEnrolledTime)
	sb.WriteString("_")
	sb.WriteString(c.Commenter)
	sb.WriteString("_")
	sb.WriteString(c.EnrolledTime)
	return sb.String()
}

//GetBoardKey ...
func (c *Comment) GetBoardKey() string {
	var sb strings.Builder
	sb.WriteString(c.BoardCompany)
	sb.WriteString("_")
	sb.WriteString(c.BoardEnrolledTime)
	return sb.String()
}

//SetDID ...
func (c *Comment) SetDID(did string) error {
	c.DID = did
	return nil
}

//SetBoardCompany ...
func (c *Comment) SetBoardCompany(boardCompany string) error {
	c.BoardCompany = boardCompany
	return nil
}

//SetBoardEnrolledTime ...
func (c *Comment) SetBoardEnrolledTime(boardEntime string) error {
	c.BoardEnrolledTime = boardEntime
	return nil
}

//SetParentCommenter ...
func (c *Comment) SetParentCommenter(parentCommenter string) error {
	c.ParentCommenter = parentCommenter
	return nil
}

//SetParentEnrolledTime ...
func (c *Comment) SetParentEnrolledTime(parentEntime string) error {
	c.ParentEnrolledTime = parentEntime
	return nil
}

//SetCommenter ...
func (c *Comment) SetCommenter(commenter string) error {
	c.Commenter = commenter
	return nil
}

//SetContent ...
func (c *Comment) SetContent(content string) error {
	c.Content = content
	return nil
}

//SetHierarchyIndex ...
func (c *Comment) SetHierarchyIndex() error {
	var sb strings.Builder
	if c.ParentCommenter == "" {
		sb.WriteString(c.EnrolledTime)
	} else {
		sb.WriteString(c.ParentEnrolledTime)
		sb.WriteString("_")
		sb.WriteString(c.EnrolledTime)
	}
	c.HierarchyIndex = sb.String()
	return nil
}

//SetEnrolledTime ...
func (c *Comment) SetEnrolledTime(entime string) error {
	c.EnrolledTime = entime
	return nil
}

//SetUpdatedTime ...
func (c *Comment) SetUpdatedTime(uptime string) error {
	c.UpdatedTime = uptime
	return nil
}
