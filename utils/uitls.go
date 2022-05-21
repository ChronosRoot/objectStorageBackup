package utils

import (
	"os"
	"path"
	"strings"
	"time"
)

func UpLoadFileKey(localFilePath string) string {
	systemHostname, _ := os.Hostname()
	FileNameAll := path.Base(localFilePath)
	FileExt := path.Ext(FileNameAll)
	FileName := strings.TrimSuffix(FileNameAll, FileExt)
	Hostname := strings.Split(systemHostname, ".")[0]
	UploadDate := time.Now().Format("2006-01-02")
	UploadFileName := FileName + "-" + UploadDate + FileExt
	return strings.Join([]string{Hostname, UploadFileName}, "/")
}
