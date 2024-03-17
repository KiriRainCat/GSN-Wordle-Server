package util

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func GetCwd() string {
	var path string
	if gin.Mode() == gin.ReleaseMode {
		path, _ = os.Executable()
	}
	return filepath.Dir(path)
}

func EnsureNessesaryDirs() {
	for _, dir := range []string{
		"./log",
	} {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}
	}
}

func Encode(obj any) []byte {
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	return bytes
}

func GetSecondsRemainForDay() time.Duration {
	todayLast := time.Now().Format("2006-01-02") + " 23:59:59"
	todayLastTime, _ := time.ParseInLocation("2006-01-02 15:04:05", todayLast, time.Local)
	return time.Duration(todayLastTime.Unix()-time.Now().Local().Unix()) * time.Second
}
