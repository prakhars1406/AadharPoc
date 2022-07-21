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
	"sync"
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

		getAadharDetailChannel := make(chan model.AadharDetails)
		//getImageDataChannel := make(chan []byte)
		var getAadharWaitGroup sync.WaitGroup
		getAadharWaitGroup.Add(1)
		go dataStoreClient.GetAadharDetails(mux.Vars(request)["id"], getAadharDetailChannel, &getAadharWaitGroup)
		//go readImageData(mux.Vars(request)["id"], getImageDataChannel, &getAadharWaitGroup)
		aadharDetails := <-getAadharDetailChannel
		imageData := config.IMAGE_BASE64
		getAadharWaitGroup.Wait()
		if aadharDetails.Error == nil {
			//fileName := "attachment; filename= " + aadharDetails.Id + ".xml"
			logrus.Info(utility.GetFuncName(), "::Get aadhar details success")
			//writer.Header().Set("Content-Disposition", fileName)
			writer.Write(getXmlData(aadharDetails, imageData))
			writer.WriteHeader(http.StatusOK)

		} else {
			writer.WriteHeader(http.StatusBadRequest)
		}
	}
}

//func readImageData(id string, getImageDataChannel chan []byte, getAadharWaitGroup *sync.WaitGroup) {
//	defer getAadharWaitGroup.Done()
//	//////Reading from file/////
//	timeout := time.Duration(5) * time.Second
//	transport := &http.Transport{
//		ResponseHeaderTimeout: timeout,
//		Dial: func(network, addr string) (net.Conn, error) {
//			return net.DialTimeout(network, addr, timeout)
//		},
//		DisableKeepAlives: true,
//	}
//	client := &http.Client{
//		Transport: transport,
//	}
//	resp, err := client.Get("https://upload.wikimedia.org/wikipedia/commons/9/9e/Placeholder_Person.jpg?20190819145659")
//	if err != nil {
//		getImageDataChannel <- []byte{}
//	}
//	// read response body
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		getImageDataChannel <- []byte{}
//	}
//	// close response body
//	defer resp.Body.Close()
//	getImageDataChannel <- body
//}

func getXmlData(aadharDetails model.AadharDetails, imageData string) []byte {
	aadharXmlData := model.AadharXmlDetails{Id: aadharDetails.Id, Name: aadharDetails.Name, PhoneNumber: aadharDetails.PhoneNumber, DateOfBirth: aadharDetails.DateOfBirth,
		Image: imageData, Signature: imageData, RightHandFingerPrint: imageData, LeftHandFingerPrint: imageData}
	data, err := xml.MarshalIndent(aadharXmlData, "", "  ")
	if err != nil {
		return []byte{}
	}
	return data
}
