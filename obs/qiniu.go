package obs

import (
	"context"
	"fmt"
	"objectStorageBackup/config"
	"objectStorageBackup/utils"
	"os"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// func
func qiniuClient() (upToken string) {
	config.SetConfig()
	accessKey := config.Config.GetString("Qiniu.accessKey")
	secretKey := config.Config.GetString("Qiniu.secretKey")
	bucket := config.Config.GetString("Qiniu.bucket")
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken = putPolicy.UploadToken(mac)
	return upToken
}

func QiniuUpLoad(localFilePath string) {
	upToken := qiniuClient()
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
