package main

import (
	"Aadhar_POC/config"
	"Aadhar_POC/database"
	"Aadhar_POC/handler"
	"Aadhar_POC/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"runtime"
)

func init() {
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", utility.FormatFilePath(f.File), f.Line)
		},
	}
	logrus.SetFormatter(formatter)
}

func main() {
	logrus.Info(utility.GetFuncName(), "::Welcome")
	router := gin.Default()
	dataStoreClient := database.MongoConnector()
	addAadharHandler := handler.AddAadharHandler(dataStoreClient)
	getAadharHandler := handler.GetAadharHandler(dataStoreClient)
	router.POST("/aadhar", addAadharHandler)
	//router.Handle("/aadhar", addAadharHandler).Methods(http.MethodPost)
	router.GET("/aadhar/:id", getAadharHandler)
	router.Run(config.KEY_SEPARATOR + config.PORT)
	//log.Fatal(http.ListenAndServe(config.KEY_SEPARATOR+config.PORT, router))

}
