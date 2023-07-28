package dto

// UserDIDUpdateRequest ...
type UserDIDUpdateRequest struct {
	Name         string 			 `json:"name"`
	PhoneNumber  string 			 `json:"phone_number"`
	DeviceToken	 string				 `json:"device_token"`
	CI 			 string				 `json:"ci"`
	ImagePath	 string				 `json:"image_path"`
	MobileOS	 string				 `json:"mobile_os"`
	CompanyName  string 			 `json:"company_name"`
	Department	 string 			 `json:"department"`
	Position     string 			 `json:"position"`
	Birthday	 string 			 `json:"birthday"`
	State        uint16 			 `json:"state"`
	EnrolledTime string				 `json:"enrolled_time"`
	BeforeDID	 string				 `json:"before_did"`
	DID			 string				 `json:"did"`
}
