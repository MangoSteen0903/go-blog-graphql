package awsLoader

import (
	"context"
	"log"
	"os"

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
		o.Region = "ap-northeast-2"
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
