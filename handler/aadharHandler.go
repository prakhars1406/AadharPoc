package handler

import (
	"Aadhar_POC/config"
	"Aadhar_POC/database"
	"Aadhar_POC/model"
	"Aadhar_POC/utility"
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PostResponse struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}


func AddAadharHandler(dataStoreClient database.MongoClient) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer utility.PanicHandler(writer, request)

		var aadharDetails model.AadharDetails

		err := json.NewDecoder(request.Body).Decode(&aadharDetails)
		if err != nil {
			writer.WriteHeader(400)
			logrus.Error(utility.GetFuncName(), "Error in POST Aadhar: ", err)
			return
		}

		var id string
		id, err = dataStoreClient.InsertAadharDetails(aadharDetails)

		if err != nil {
			writer.WriteHeader(500)
			logrus.Error(utility.GetFuncName(), "Error in POST Aadhar: ", err)
			return
		}

		person := PostResponse{Id: id, Message: "User created successfully"}

		jsonResponse, _ := json.Marshal(person)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(jsonResponse)
	}
}

func GetAadharHandler(dataStoreClient database.MongoClient) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer utility.PanicHandler(writer, request)
		aadharDetails,err:=dataStoreClient.GetAadharDetails(mux.Vars(request)["id"])
		imageData := config.IMAGE_BASE64
		if err == nil {
			logrus.Info(utility.GetFuncName(), "::Get aadhar details success")
			writer.Header().Set("Content-Type", "application/xml")
			writer.Write(getXmlData(aadharDetails, imageData))

		} else {
			logrus.Info("::Get aadhar details failed with:", err, mux.Vars(request)["id"])
			writer.WriteHeader(http.StatusBadRequest)
		}
	}
}

func getXmlData(aadharDetails map[string]string, imageData string) []byte {
	aadharXmlData := model.AadharXmlDetails{Id: aadharDetails["id"], Name: aadharDetails["name"], PhoneNumber: aadharDetails["phonenumber"], DateOfBirth: aadharDetails["dateofbirth"],
		AddressLine1:aadharDetails["addressline1"],AddressLine2:aadharDetails["addressline2"],Pincode: aadharDetails["pincode"],City: aadharDetails["city"],State: aadharDetails["state"],
		Image:imageData, Signature: imageData}
	data, err := xml.MarshalIndent(aadharXmlData, "", "  ")
	if err != nil {
		return []byte{}
	}
	return data
}
