package images

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	ACCOUNT_ID        = os.Getenv("ACCOUNT_ID")
	ACCESS_KEY_ID     = os.Getenv("ACCESS_KEY_ID")
	ACCESS_KEY_SECRET = os.Getenv("ACCESS_KEY_SECRET")
	ENDPOINT          = os.Getenv("ENDPOINT") // r2.cloudflarestorage.com
)

var BUCKET_NAME = os.Getenv("BUCKET_NAME")

var App *S3App

type S3App struct {
	Client     *s3.Client
	BucketName string
}

func NewClient(accountId string, accessKeyId string, accessKeySecret string, endpoint string) (*s3.Client, error) {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{URL: fmt.Sprintf("https://%s.%s", accountId, endpoint)}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyId, accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(cfg), nil
}

func InitS3App() error {
	client, err := NewClient(ACCOUNT_ID, ACCESS_KEY_ID, ACCESS_KEY_SECRET, ENDPOINT)
	if err != nil {
		return err
	}

	App = &S3App{
		Client:     client,
		BucketName: BUCKET_NAME,
	}

	return nil
}
