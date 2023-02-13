package model

import (
	"context"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"qiublog/utils"
	"qiublog/utils/errmsg"
)

var Oss = utils.Config.Server.Oss

func UpLoadFile(uploadName string, file multipart.File, fileSize int64) (int, string) {

	if Oss == "qiniu" {
		return uploadQiniu(uploadName, file, fileSize)
	} else if Oss == "aliyun" {
		return uploadAliyun(uploadName, file, fileSize)
	} else {
		return errmsg.ERROR, ""
	}
}

func uploadQiniu(uploadName string, file multipart.File, fileSize int64) (int, string) {
	var AccessKey = utils.Config.Oss.QiniuAccessKey
	var SecretKey = utils.Config.Oss.QiniuSecretKey
	var Bucket = utils.Config.Oss.QiniuBucket
	var ImgUrl = utils.Config.Oss.QiniuSever
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	region, _ := storage.GetRegion(AccessKey, Bucket)
	cfg := storage.Config{
		Region:        region,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err := formUploader.Put(context.Background(), &ret, upToken, uploadName, file, fileSize, &putExtra)
	if err != nil {
		return errmsg.ERROR, ""
	}
	url := ImgUrl + ret.Key
	return errmsg.SUCCESS, url
}

func uploadAliyun(uploadName string, file multipart.File, fileSize int64) (int, string) {
	var AliyunAccessKeyId = utils.Config.Oss.AliyunAccessKeyId
	var AliyunAccessKeySecret = utils.Config.Oss.AliyunAccessKeySecret
	var AliyunEndpoint = utils.Config.Oss.AliyunEndpoint
	var AliyunBucketName = utils.Config.Oss.AliyunBucketName
	client, err := oss.New(AliyunEndpoint, AliyunAccessKeyId, AliyunAccessKeySecret)
	if err != nil {
		return errmsg.ERROR, ""
	}
	// 获取存储空间。
	bucket, err := client.Bucket(AliyunBucketName)
	if err != nil {
		return errmsg.ERROR, ""
	}
	err = bucket.PutObject(uploadName, file)
	if err != nil {
		return errmsg.ERROR, ""
	}
	//拼接链接,默认使用https
	return errmsg.SUCCESS, fmt.Sprintf("https://%s.%s/%s", AliyunBucketName, AliyunEndpoint, uploadName)
}
