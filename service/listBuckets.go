package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	ACCESS_KEY = "FIXME"
	SECRET_KEY = "FIXME"
	ENDPOINT   = "FIXME"
)

func initS3Client() *s3.S3 {
	// 构建配置项
	cfg := &aws.Config{
		Region:                        aws.String("us-east-1"), //By default, Don't edit it
		Credentials:                   credentials.NewStaticCredentials(ACCESS_KEY, SECRET_KEY, ""),
		Endpoint:                      aws.String(ENDPOINT),
		S3DisableContentMD5Validation: aws.Bool(true),
		S3ForcePathStyle:              aws.Bool(true),
	}

	// 实例化 Session 对象
	sess := session.Must(session.NewSession(cfg))

	// 实例化 Service Client
	svc := s3.New(sess)

	return svc
}

func main() {
	svc := initS3Client()
	input := &s3.ListBucketsInput{}

	result, err := svc.ListBuckets(input)
	if err != nil {
		if aErr, ok := err.(awserr.Error); ok {
			switch aErr.Code() {
			default:
				fmt.Println(aErr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
	return
}
