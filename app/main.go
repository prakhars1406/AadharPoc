package main

import (
	"Aadhar_POC/config"
	"Aadhar_POC/handler"
	"Aadhar_POC/utility"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
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
	router := mux.NewRouter().StrictSlash(true)
	// dataStoreClient := database.MongoConnector()
	// dataStoreClient := database.LevelDBConnector()
	addAadharHandler := handler.AddAadharHandler()
	getAadharHandler := handler.GetAadharHandler()
	router.Handle("/aadhar", addAadharHandler).Methods(http.MethodPost)
	router.Handle("/aadhar/{id}", getAadharHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(config.KEY_SEPARATOR+config.PORT, router))

}
