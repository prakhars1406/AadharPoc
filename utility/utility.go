package utility

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"runtime/debug"
	"strings"
)

func FormatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func GetFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	i := strings.LastIndex(f.Name(), "/")
	/*i2 := strings.LastIndex(f.Name(), "func1")
	if i2==-1{
		return f.Name()[i+1:]
	}else{
		return f.Name()[i+1:i2-1]
	}*/
	return f.Name()[i+1:]
}

func PanicHandler(w http.ResponseWriter, r *http.Request) {
	if r := recover(); r != nil {
		log.Println(fmt.Sprintf("Recovered in f %v", r))
		var err error
		switch x := r.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = errors.New("Unknown panic")
		}
		if err != nil {
			fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func WriteAadharDetailsToXMLFile(b []byte, ID string) bool {
	err := ioutil.WriteFile(ID+".xml", b, 0644)
	if err != nil {
		return false
	}
	return true
}
