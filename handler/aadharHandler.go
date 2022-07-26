package handler

import (
	"Aadhar_POC/database"
	"Aadhar_POC/model"
	"Aadhar_POC/utility"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type PostResponse struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func AddAadharHandler(dataStoreClient database.MongoClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer utility.PanicHandler(ctx)

		var aadharDetails model.AadharDetails

		err := json.NewDecoder(ctx.Request.Body).Decode(&aadharDetails)
		if err != nil {
			ctx.Writer.WriteHeader(400)
			logrus.Error(utility.GetFuncName(), "Error in POST Aadhar: ", err)
			return
		}

		var id string
		id, err = dataStoreClient.InsertAadharDetails(aadharDetails)

		if err != nil {
			ctx.Writer.WriteHeader(500)
			logrus.Error(utility.GetFuncName(), "Error in POST Aadhar: ", err)
			return
		}

		person := PostResponse{Id: id, Message: "User created successfully"}

		jsonResponse, _ := json.Marshal(person)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusOK)
		_, _ = ctx.Writer.Write(jsonResponse)
	}
}

func GetAadharHandler(dataStoreClient database.MongoClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer utility.PanicHandler(ctx)
		aadharDetails, err := dataStoreClient.GetAadharDetails(ctx.Param("id"))
		if err == nil {

			logrus.Info("::Get aadhar details success")
			ctx.Writer.Header().Set("Content-Disposition", "test.xml")
			ctx.Writer.Write(aadharDetails)
			ctx.Writer.WriteHeader(http.StatusOK)

		} else {
			logrus.Error("::Get aadhar details failed with:", err, ctx.Param("id"))
			ctx.Writer.WriteHeader(http.StatusBadRequest)
		}
	}
}
