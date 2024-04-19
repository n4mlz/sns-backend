package s3

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
	ACCOUNT_ID        = os.Getenv("S3_ACCOUNT_ID")
	ACCESS_KEY_ID     = os.Getenv("S3_ACCESS_KEY_ID")
	ACCESS_KEY_SECRET = os.Getenv("S3_ACCESS_KEY_SECRET")
	ENDPOINT          = os.Getenv("S3_ENDPOINT")
)

var BUCKET_NAME = os.Getenv("S3_BUCKET_NAME")

type S3App struct {
	Client     *s3.Client
	BucketName string
}

func newClient(accountId string, accessKeyId string, accessKeySecret string, endpoint string) (*s3.Client, error) {
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

func NewS3App() (*S3App, error) {
	client, err := newClient(ACCOUNT_ID, ACCESS_KEY_ID, ACCESS_KEY_SECRET, ENDPOINT)
	if err != nil {
		return nil, err
	}

	return &S3App{
		Client:     client,
		BucketName: BUCKET_NAME,
	}, nil
}
