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
	ACCESS_KEY  = "FIXME"
	SECRET_KEY  = "FIXME"
	BUCKET_NAME = "FIXME"
	ENDPOINT    = "FIXME"
)

func initS3Client() *s3.S3 {
	// initial config
	cfg := &aws.Config{
		Region:                        aws.String("us-east-1"), //By default, Don't edit it
		Credentials:                   credentials.NewStaticCredentials(ACCESS_KEY, SECRET_KEY, ""),
		Endpoint:                      aws.String(ENDPOINT),
		S3DisableContentMD5Validation: aws.Bool(true),
		S3ForcePathStyle:              aws.Bool(true),
	}

	// new session
	sess := session.Must(session.NewSession(cfg))

	// new svc
	svc := s3.New(sess)

	return svc
}

func listMulitpartUploads(svc *s3.S3) {
	input := &s3.ListMultipartUploadsInput{
		Bucket: aws.String(BUCKET_NAME),
	}

	result, err := svc.ListMultipartUploads(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
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

func main() {
	svc := initS3Client()
	listMulitpartUploads(svc)
}
