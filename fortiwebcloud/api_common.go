package fortiwebcloud

import (
	"os"
	"time"
)

type APICommonInterface interface {
	NewRequest(d interface{}) *Request
	Send() error
	ReadData() (interface{}, error)
}

func FileExist(filename string) bool {

	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func output(v string) {

	if !FileExist("/var/log/cloudwaf_debug") {
		return
	}
	f, err := os.OpenFile("output.debug", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		panic(err)
	}
	fd_time := time.Now().Format("2006-01-02 15:04:05 ::")
	defer f.Close()
	f.WriteString(fd_time + v + "\n")
	return
}
func OutPut(v string) {
	output(v)
	return
}
