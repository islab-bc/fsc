package dto

// ChaincodeResponse is response for chaincode
type ChaincodeResponse struct {
	ChaincodeResult string `json:"data"`
	Status          string `json:"message"`
	StatusCode      int    `json:"status_code"`
}
