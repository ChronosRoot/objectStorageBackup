package main

import (
	"objectStorageBackup/obs"
	"os"
)

func main() {
	for _, localFilePath := range os.Args[1:] {
		obs.QiniuUpLoad(localFilePath)
	}
}
