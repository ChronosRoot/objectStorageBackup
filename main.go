package main

import (
	"objectStorageBackup/config"
	"objectStorageBackup/qiniu"
	"os"
)

func main() {
	config.SetConfig()

	accessKey := config.Config.GetString("Qiniu.accessKey")
	secretKey := config.Config.GetString("Qiniu.secretKey")
	bucket := config.Config.GetString("Qiniu.bucket")
	for _, localFilePath := range os.Args[1:] {
		qiniu.UpLoadQiniu(accessKey, secretKey, bucket, localFilePath)
	}
}
