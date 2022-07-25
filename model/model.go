package model

type AadharDetails struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	PhoneNumber  string `json:"phoneNumber,omitempty"`
	DateOfBirth  string `json:"dateOfBirth,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressLine2 string `json:"addressLine2,omitempty"`
	Pincode      string `json:"pincode,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
}
