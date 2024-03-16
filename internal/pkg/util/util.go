package util

import (
	"encoding/json"
	"os"
	"path/filepath"

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
