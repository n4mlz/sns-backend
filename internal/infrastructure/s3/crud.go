package s3

import (
	"bytes"
	"context"
	"io"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
)

func userIconImageUrl(user *userDomain.User) string {
	return path.Join("images", "users", user.UserName.String(), "icon.png")
}

func userBgImageUrl(user *userDomain.User) string {
	return path.Join("images", "users", user.UserName.String(), "background.png")
}

func (app *S3App) saveObject(objectKey string, object []byte) error {
	_, err := app.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(app.BucketName),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(object),
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) deleteObject(objectKey string) error {
	_, err := app.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(app.BucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) moveObject(sourceObjectKey string, targetObjectKey string) error {
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

	err = app.saveObject(targetObjectKey, buf)
	if err != nil {
		return err
	}

	err = app.deleteObject(sourceObjectKey)
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) SaveIcon(user *userDomain.User, file io.Reader) error {
	fileBytes, err := fotmatImageForIcon(file)
	if err != nil {
		return err
	}

	objectKey := userIconImageUrl(user)

	err = app.saveObject(objectKey, fileBytes)
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) SaveBgImage(user *userDomain.User, file io.Reader) error {
	fileBytes, err := fotmatImageForBgImage(file)
	if err != nil {
		return err
	}

	objectKey := userBgImageUrl(user)

	err = app.saveObject(objectKey, fileBytes)
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) DeleteIcon(user *userDomain.User) error {
	objectKey := userIconImageUrl(user)

	if err := app.deleteObject(objectKey); err != nil {
		return err
	}
	return nil
}

func (app *S3App) DeleteBgImage(user *userDomain.User) error {
	objectKey := userBgImageUrl(user)

	if err := app.deleteObject(objectKey); err != nil {
		return err
	}
	return nil
}

func (app *S3App) MoveResources(sourceUser *userDomain.User, targetUser *userDomain.User) error {
	err := app.moveObject(userIconImageUrl(sourceUser), userIconImageUrl(targetUser))
	if err != nil {
		return err
	}

	err = app.moveObject(userBgImageUrl(sourceUser), userBgImageUrl(targetUser))
	if err != nil {
		return err
	}

	return nil
}
