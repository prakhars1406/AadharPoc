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
	Image        []byte `json:"image,omitempty"`
	Error        error  `json:"error,omitempty"`
}
type AadharDetailsMongoResponse struct {
	Id           string `bson:"id,omitempty"`
	Name         string `bson:"name,omitempty"`
	PhoneNumber  string `bson:"phoneNumber,omitempty"`
	DateOfBirth  string `bson:"dateOfBirth,omitempty"`
	AddressLine1 string `bson:"addressLine1,omitempty"`
	AddressLine2 string `bson:"addressLine2,omitempty"`
	Pincode      string `bson:"pincode,omitempty"`
	City         string `bson:"city,omitempty"`
	State        string `bson:"state,omitempty"`
	Image        []byte `bson:"image,omitempty"`
	Error        error  `bson:"error,omitempty"`
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
	Image                []byte `xml:"image,omitempty"`
	Signature            []byte `xml:"signature,omitempty"`
	RightHandFingerPrint []byte `xml:"rightHandFingerPrint,omitempty"`
	LeftHandFingerPrint  []byte `xml:"leftHandFingerPrint,omitempty"`
	Error                error  `xml:"error,omitempty"`
}
