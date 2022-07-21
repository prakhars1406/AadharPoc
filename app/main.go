package main

import (
	"Aadhar_POC/config"
	"Aadhar_POC/database"
	"Aadhar_POC/handler"
	"Aadhar_POC/utility"
	"compress/gzip"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
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
// Gzip Compression
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func Gzip(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			handler.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		handler.ServeHTTP(gzw, r)
	})
}

func main()  {
	logrus.Info(utility.GetFuncName(), "::Welcome")
	router := mux.NewRouter().StrictSlash(true)
	dataStoreClient := database.MongoConnector()
	addAadharHandler := handler.AddAadharHandler(dataStoreClient)
	getAadharHandler := Gzip(handler.GetAadharHandler(dataStoreClient))
	router.Handle("/aadhar", addAadharHandler).Methods(http.MethodPost)
	router.Handle("/aadhar/{id}", getAadharHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(config.KEY_SEPARATOR+config.PORT, router))

}
