package utils

import (
	"os"
	"path"
	"strings"
	"time"
)

func UpLoadFileKey(localFilePath string) (UpLoadFileKey string) {
	systemHostname, _ := os.Hostname()
	FileNameAll := path.Base(localFilePath)
	Hostname := strings.Split(systemHostname, ".")[0]
	UploadDate := time.Now().Format("2006-01-02")
	UploadFileName := UploadDate + "-" + FileNameAll
	UpLoadFileKey = strings.Join([]string{Hostname, UploadFileName}, "/")
	return UpLoadFileKey
}
