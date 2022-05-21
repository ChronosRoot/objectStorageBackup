package qiniu

import (
	"context"
	"fmt"
	"objectStorageBackup/utils"
	"os"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// func
func OBSClient(accessKey string, secretKey string, bucket string) (upToken string) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken = putPolicy.UploadToken(mac)
	return upToken
}

func UpLoadQiniu(accessKey, secretKey, bucket, localFilePath string) {
	upToken := OBSClient(accessKey, secretKey, bucket)
	key := utils.UpLoadFileKey(localFilePath)

	cfg := storage.Config{}
	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	recorder, err := storage.NewFileRecorder(os.TempDir())
	if err != nil {
		fmt.Println(err)
		return
	}
	putExtra := storage.RputV2Extra{
		Recorder: recorder,
	}
	err = resumeUploader.PutFile(context.Background(), &ret, upToken, key, localFilePath, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}
