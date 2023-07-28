package dto

// UserGetByQueryRequest ...
type UserGetByQueryRequest struct {
	QueryString string `json:"querystring"`
}

// UserGetByUserRequest ...
type UserGetByUserRequest struct {
	DID 		string `json:"did"`
}

// UserGetByPhoneRequest ...
type UserGetByPhoneRequest struct {
	PhoneNumber string `json:"phone_number"`
}

// UserGetByUpdateRequest ...
type UserGetByUpdateRequest struct {
	CI				string `json:"ci"`
}

// AdmissionGetByUserRequest ...
type AdmissionGetByUserRequest struct {
	DID 		string `json:"did"`
}