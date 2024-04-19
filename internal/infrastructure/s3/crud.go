package s3

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (app *S3App) SaveImage(objectKey string, file *os.File) error {
	fileBytes, err := fotmatImage(file)
	if err != nil {
		return err
	}

	err = app.SaveBinary(objectKey, fileBytes)
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) SaveBinary(objectKey string, fileBytes []byte) error {
	_, err := app.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(app.BucketName),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(fileBytes),
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) Delete(objectKey string) error {
	_, err := app.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(app.BucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) Move(sourceObjectKey string, targetObjectKey string) error {
	result, err := app.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(app.BucketName),
		Key:    aws.String(sourceObjectKey),
	})
	if err != nil {
		// Object not found
		return nil
	}
	defer result.Body.Close()

	buf, err := io.ReadAll(result.Body)
	if err != nil {
		return err
	}

	err = app.SaveBinary(targetObjectKey, buf)
	if err != nil {
		return err
	}

	err = app.Delete(sourceObjectKey)
	if err != nil {
		return err
	}

	return nil
}
