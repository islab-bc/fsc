package model

// User ...
type User struct {
	DocType      string `json:"doctype"`
	Name         string `json:"name"`
	PhoneNumber  string `json:"phone_number"`
	DeviceToken  string `json:"device_token"`
	MobileOS     string `json:"mobile_os"`
	CI           string `json:"ci"`
	ImagePath    string `json:"image_path"`
	CompanyName  string `json:"company_name"`
	Department   string `json:"department"`
	Position     string `json:"position"`
	Birthday     string `json:"birthday"`
	State        uint16 `json:"state"`
	EnrolledTime string `json:"enrolled_time"`
	Location     string `json:"location"`
}

// NewUser ...
func NewUser() *User {
	return &User{
		DocType: "user",
	}
}

// GetDocType ...
func (u *User) GetDocType() string {
	return u.DocType
}

// GetName ...
func (u *User) GetName() string {
	return u.Name
}

// GetPhoneNumber ...
func (u *User) GetPhoneNumber() string {
	return u.PhoneNumber
}

// GetDeviceToken ...
func (u *User) GetDeviceToken() string {
	return u.DeviceToken
}

// GetMobileOS ...
func (u *User) GetMobileOS() string {
	return u.MobileOS
}

// GetCI ...
func (u *User) GetCI() string {
	return u.CI
}

// GetImagePath ...
func (u *User) GetImagePath() string {
	return u.ImagePath
}

// GetCompanyName ...
func (u *User) GetCompanyName() string {
	return u.CompanyName
}

// GetDepartment ...
func (u *User) GetDepartment() string {
	return u.Department
}

// GetPosition ...
func (u *User) GetPosition() string {
	return u.Position
}

// GetBirthday ...
func (u *User) GetBirthday() string {
	return u.Birthday
}

// GetState ...
func (u *User) GetState() uint16 {
	return u.State
}

// GetEnrolledTime ...
func (u *User) GetEnrolledTime() string {
	return u.EnrolledTime
}

// GetLocation ...
func (u *User) GetLocation() string {
	return u.Location
}

// SetName ...
func (u *User) SetName(name string) error {
	u.Name = name
	return nil
}

// SetPhoneNumber ...
func (u *User) SetPhoneNumber(pnum string) error {
	u.PhoneNumber = pnum
	return nil
}

// SetDeviceToken ...
func (u *User) SetDeviceToken(token string) error {
	u.DeviceToken = token
	return nil
}

// SetMobileOS ...
func (u *User) SetMobileOS(os string) error {
	u.MobileOS = os
	return nil
}

// SetCI ...
func (u *User) SetCI(ci string) error {
	u.CI = ci
	return nil
}

// SetCI ...
func (u *User) SetImagePath(path string) error {
	u.ImagePath = path
	return nil
}

// SetCompanyName ...
func (u *User) SetCompanyName(cname string) error {
	u.CompanyName = cname
	return nil
}

// SetDepartment ...
func (u *User) SetDepartment(department string) error {
	u.Department = department
	return nil
}

// SetPosition ...
func (u *User) SetPosition(postion string) error {
	u.Position = postion
	return nil
}

// SetBirthday ...
func (u *User) SetBirthday(birthday string) error {
	u.Birthday = birthday
	return nil
}

// SetState ...
func (u *User) SetState(state uint16) error {
	u.State = state
	return nil
}

// SetEnrolledTime ...
func (u *User) SetEnrolledTime(entime string) error {
	u.EnrolledTime = entime
	return nil
}

// SetLocation ...
func (u *User) SetLocation(location string) error {
	u.Location = location
	return nil
}
