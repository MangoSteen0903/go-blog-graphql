package awsLoader

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func LoadAWS() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(os.Getenv("S3_ACCESS_KEY"), os.Getenv("S3_SECRET_KEY"), ""),
		))
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Region = os.Getenv("S3_REGION")
	})

	return client
	/*
		resp, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

		var buckets []types.Bucket
		if err != nil {
			log.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
		} else {
			buckets = resp.Buckets
		}

		fmt.Println(buckets)
	*/
}

func UploadFile(dirname string, username string, file *graphql.Upload) (*string, error) {

	time := fmt.Sprintf("%v-%v-%v", time.Now().Year(), time.Now().Day(), time.Now().Hour())
	userFileName := fmt.Sprintf("%v/%v-%v-%v.jpg", dirname, username, file.Filename, time)

	client := LoadAWS()
	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(userFileName),
		Body:   file.File,
	})
	if err != nil {
		return nil, err
	}
	url := "https://%s.s3.%s.amazonaws.com/%s"
	url = fmt.Sprintf(url, os.Getenv("BUCKET_NAME"), os.Getenv("S3_REGION"), userFileName)

	return &url, nil
}
