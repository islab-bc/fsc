package model

import "strings"

// Admission ...
type Admission struct {
	DocType      string 			 `json:"doctype"`
	Location     string 			 `json:"location"`
	DID			 string				 `json:"did"`
	EnrolledTime string				 `json:"enrolled_time"`
	UpdatedTime	 string				 `json:"updated_time"`
}

// NewAdmission ...
func NewAdmission() *Admission {
	return &Admission{
		DocType:     "admission",
	}
}

// GetDocType ...
func (a *Admission) GetDocType() string {
	return a.DocType
}

// GetLocation ...
func (a *Admission) GetLocation() string {
	return a.Location
}

// GetDID ...
func (a *Admission) GetDID() string {
	return a.DID
}

// GetEnrolledTime ...
func (a *Admission) GetEnrolledTime() string {
	return a.EnrolledTime
}

// GetUpdatedTime ...
func (a *Admission) GetUpdatedTime() string {
	return a.UpdatedTime
}

//GetKey ...
func (a *Admission) GetKey() string {
	var sb strings.Builder
	sb.WriteString(a.DID)
	sb.WriteString("_")
	sb.WriteString(a.Location)
	sb.WriteString("_")
	sb.WriteString(a.EnrolledTime)
	return sb.String()
}

// SetLocation ...
func (a *Admission) SetLocation(location string) error {
	a.Location = location
	return nil
}

// SetDID ...
func (a *Admission) SetDID(did string) error {
	a.DID = did
	return nil
}

// SetEnrolledTime ...
func (a *Admission) SetEnrolledTime(time string) error {
	a.EnrolledTime = time
	return nil
}

// SetUpdatedTime ...
func (a *Admission) SetUpdatedTime(time string) error {
	a.UpdatedTime = time
	return nil
}