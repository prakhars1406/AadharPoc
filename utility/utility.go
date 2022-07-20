package utility

import (
	"runtime"
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
