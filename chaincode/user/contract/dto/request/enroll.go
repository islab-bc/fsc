package dto

// UserEnrollRequest ...
type UserEnrollRequest struct {
	Name         string 			 `json:"name"`
	PhoneNumber  string 			 `json:"phone_number"`
	DeviceToken	 string				 `json:"device_token"`
	CI 			 string				 `json:"ci"`
	DID			 string				 `json:"did"`
	MobileOS	 string				 `json:"mobile_os"`
	ImagePath	 string				 `json:"image_path"`
	CompanyName  string 			 `json:"company_name"`
	Department	 string 			 `json:"department"`
	Position     string 			 `json:"position"`
	Birthday	 string 			 `json:"birthday"`
	State        uint16 			 `json:"state"`
	EnrolledTime string 			 `json:"enrolled_time"`
}

// AdmissionEnrollRequest ...
type AdmissionEnrollRequest struct {
	DID			 string				 `json:"did"`
	Location	 string				 `json:"location"`
	EnrolledTime string				 `json:"enrolled_time"`
}