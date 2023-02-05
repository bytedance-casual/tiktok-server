package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"tiktok-server/internal/conf"
)

var (
	client *oss.Client
	bucket *oss.Bucket
	baseURL string
)

func Init() {
	ossConfig := conf.Config.OssAliyun
	baseURL = ossConfig.BaseURL
	// 创建OSSClient实例
	client, err := oss.New(ossConfig.Endpoint, ossConfig.AccessKeyID, ossConfig.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	// 获取存储空间
	bucket, err = client.Bucket(ossConfig.BucketName)
	if err != nil {
		panic(err)
	}
}
