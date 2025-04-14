package model

import (
	"Blog/utils"
	"Blog/utils/errormsg"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// 自定义返回值结构体

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiniuServer

func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	region, err1 := storage.GetRegion(AccessKey, Bucket)
	if err1 != nil {
		fmt.Println("获取区域失败：", err1)
	return "", errormsg.ERROR
	}
	cfg := storage.Config{
		Region:        region,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		fmt.Println("上传失败：", err)
		return "1", errormsg.ERROR
	}
	url := ImgUrl + ret.Key
	return url, errormsg.SUCCSE

}
