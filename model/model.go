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
type AadharXmlDetails struct {
	Id                   string `xml:"id,omitempty"`
	Name                 string `xml:"name,omitempty"`
	PhoneNumber          string `xml:"phoneNumber,omitempty"`
	DateOfBirth          string `xml:"dateOfBirth,omitempty"`
	AddressLine1         string `xml:"addressLine1,omitempty"`
	AddressLine2         string `xml:"addressLine2,omitempty"`
	Pincode              string `xml:"pincode,omitempty"`
	City                 string `xml:"city,omitempty"`
	State                string `xml:"state,omitempty"`
	Image                string `xml:"image,omitempty"`
	Signature            string `xml:"signature,omitempty"`
}
