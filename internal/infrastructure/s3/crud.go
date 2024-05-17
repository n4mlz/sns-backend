package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"path"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
	"github.com/rs/xid"
)

func generateIconObjectKey(user *userDomain.User) string {
	return path.Join("images", "users", user.UserName.String(), fmt.Sprintf("icon_%s.png", xid.New().String()))
}

func generateBgImageObjectKey(user *userDomain.User) string {
	return path.Join("images", "users", user.UserName.String(), fmt.Sprintf("background_%s.png", xid.New().String()))
}

func objectKeyToUrl(objectKey string) string {
	url, err := url.JoinPath(RESOURCE_URL, objectKey)
	if err != nil {
		return ""
	}
	return url
}

func urlToObjectKey(Url string) string {
	u, err := url.Parse(Url)
	if err != nil {
		return ""
	}

	objectKey := u.Path
	return objectKey[1:]
}

func (app *S3App) saveObject(objectKey string, object []byte, ContentType string) error {
	_, err := app.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(app.BucketName),
		Key:         aws.String(objectKey),
		Body:        bytes.NewReader(object),
		ContentType: aws.String(ContentType),
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

	err = app.saveObject(targetObjectKey, buf, *result.ContentType)
	if err != nil {
		return err
	}

	err = app.deleteObject(sourceObjectKey)
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) SaveIcon(user *userDomain.User, file io.Reader) (userDomain.ImageUrl, error) {
	fileBytes, err := fotmatImageForIcon(file)
	if err != nil {
		return "", err
	}

	err = app.DeleteIcon(user)
	if err != nil {
		return "", err
	}

	objectKey := generateIconObjectKey(user)

	err = app.saveObject(objectKey, fileBytes, "image/png")
	if err != nil {
		return "", err
	}

	return userDomain.ImageUrl(objectKeyToUrl(objectKey)), nil
}

func (app *S3App) SaveBgImage(user *userDomain.User, file io.Reader) (userDomain.ImageUrl, error) {
	fileBytes, err := fotmatImageForBgImage(file)
	if err != nil {
		return "", err
	}

	err = app.DeleteBgImage(user)
	if err != nil {
		return "", err
	}

	objectKey := generateBgImageObjectKey(user)

	err = app.saveObject(objectKey, fileBytes, "image/png")
	if err != nil {
		return "", err
	}

	return userDomain.ImageUrl(objectKeyToUrl(objectKey)), nil
}

func (app *S3App) DeleteIcon(user *userDomain.User) error {
	if user.IconUrl == "" {
		return nil
	}

	objectKey := urlToObjectKey(user.IconUrl.String())

	if err := app.deleteObject(objectKey); err != nil {
		return err
	}
	return nil
}

func (app *S3App) DeleteBgImage(user *userDomain.User) error {
	if user.BgImageUrl == "" {
		return nil
	}

	objectKey := urlToObjectKey(user.BgImageUrl.String())

	if err := app.deleteObject(objectKey); err != nil {
		return err
	}
	return nil
}

func (app *S3App) MoveResources(sourceUser *userDomain.User, targetUser *userDomain.User) (userDomain.ImageUrl, userDomain.ImageUrl, error) {
	iconObjectKey := generateIconObjectKey(targetUser)
	bgImageObjectKey := generateBgImageObjectKey(targetUser)

	err := app.moveObject(urlToObjectKey(sourceUser.IconUrl.String()), iconObjectKey)
	if err != nil {
		return "", "", err
	}

	err = app.moveObject(urlToObjectKey(sourceUser.BgImageUrl.String()), bgImageObjectKey)
	if err != nil {
		return "", "", err
	}

	return userDomain.ImageUrl(objectKeyToUrl(iconObjectKey)), userDomain.ImageUrl(objectKeyToUrl(bgImageObjectKey)), nil
}
