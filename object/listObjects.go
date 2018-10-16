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
	ACCESS_KEY = "liul"
	SECRET_KEY = "liul123"
	BUCKET_NAME = "software"
	ENDPOINT   = "https://oss-cn-shanghai.speedycloud.org"
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

	// new svc
	svc := s3.New(sess)

	return svc
}

func listObject(svc *s3.S3) {
	input := &s3.ListObjectsInput{
		Bucket:  aws.String(BUCKET_NAME),
		MaxKeys: aws.Int64(2),
	}

	result, err := svc.ListObjects(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

func main(){
	svc := initS3Client()
	listObject(svc)
}
