package handler

import (
	"Aadhar_POC/database"
	"fmt"
	"io"
	"net"
	"Aadhar_POC/model"
	"Aadhar_POC/utility"
	"encoding/json"
	"log"
	"net/http"
	"time"
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
			log.Println("Error in POST Aadhar: ", err)
			return
		}

		var id string
		id, err = dataStoreClient.InsertAadharDetails(aadharDetails)

		if err != nil {
			writer.WriteHeader(500)
			log.Println("Error in POST Aadhar: ", err)
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
		timeout := time.Duration(5) * time.Second
		transport := &http.Transport{
			ResponseHeaderTimeout: timeout,
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, timeout)
			},
			DisableKeepAlives: true,
		}

		client := &http.Client{
			Transport: transport,
		}
		resp, err := client.Get("https://file.io/aNuhLF8xQHV1")
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		writer.Header().Set("Content-Type", "application/octet-stream")
		writer.Header().Set("Content-Disposition", "attachment; filename=text.xml")
		io.Copy(writer, resp.Body)
	}
}
