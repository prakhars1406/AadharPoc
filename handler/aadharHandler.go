package handler

import (
	"Aadhar_POC/database"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func AddAadharHandler(dataStoreClient database.MongoClient) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
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
