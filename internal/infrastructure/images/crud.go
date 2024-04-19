package images

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (app *S3App) UploadImage(ctx context.Context, objectKey string, file *os.File) error {
	fileBytes, err := fotmatImage(file)
	if err != nil {
		return err
	}

	err = app.UploadBinary(ctx, objectKey, fileBytes)
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) UploadBinary(ctx context.Context, objectKey string, fileBytes []byte) error {
	_, err := App.Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(App.BucketName),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(fileBytes),
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) DeleteObject(ctx context.Context, objectKey string) error {
	_, err := App.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(App.BucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) MoveObject(ctx context.Context, sourceObjectKey string, targetObjectKey string) error {
	result, err := App.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(App.BucketName),
		Key:    aws.String(sourceObjectKey),
	})
	if err != nil {
		return err
	}
	defer result.Body.Close()

	buf, err := io.ReadAll(result.Body)
	if err != nil {
		return err
	}

	err = app.UploadBinary(ctx, targetObjectKey, buf)
	if err != nil {
		return err
	}

	err = app.DeleteObject(ctx, sourceObjectKey)
	if err != nil {
		return err
	}

	return nil
}
