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
	FileName := strings.SplitN(FileNameAll, ".", 2)[0]
	FileExt := strings.SplitN(FileNameAll, ".", 2)[1]
	Hostname := strings.Split(systemHostname, ".")[0]
	UploadDate := time.Now().Format("2006-01-02")
	UploadFileName := FileName + "-" + UploadDate + "." + FileExt
	return strings.Join([]string{Hostname, UploadFileName}, "/")
}
