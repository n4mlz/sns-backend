package s3

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"path"
	"strings"

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
	resource, err := url.Parse(RESOURCE_URL)
	if err != nil || resource.Scheme == "" || resource.Host == "" || u.Scheme != resource.Scheme || u.Host != resource.Host {
		return ""
	}
	objectKey := strings.TrimPrefix(u.Path, "/")
	if objectKey == "" || path.Clean(objectKey) != objectKey || !strings.HasPrefix(objectKey, "images/users/") {
		return ""
	}
	return objectKey
}

func (app *S3App) saveObject(objectKey string, object []byte, ContentType string) error {
	ctx, cancel := s3OperationContext()
	defer cancel()
	_, err := app.Client.PutObject(ctx, &s3.PutObjectInput{
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
	ctx, cancel := s3OperationContext()
	defer cancel()
	_, err := app.Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(app.BucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return err
	}

	return nil
}

func (app *S3App) moveObject(sourceObjectKey string, targetObjectKey string) error {
	ctx, cancel := s3OperationContext()
	defer cancel()
	result, err := app.Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(app.BucketName),
		Key:    aws.String(sourceObjectKey),
	})
	if err != nil {
		// Object not found
		return nil
	}
	defer result.Body.Close()

	buf, err := io.ReadAll(io.LimitReader(result.Body, maxImageBytes+1))
	if err != nil {
		return err
	}
	if len(buf) > maxImageBytes {
		return fmt.Errorf("object is too large")
	}

	contentType := "image/png"
	if result.ContentType != nil && *result.ContentType != "" {
		contentType = *result.ContentType
	}
	err = app.saveObject(targetObjectKey, buf, contentType)
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
	if objectKey == "" {
		return nil
	}

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
	if objectKey == "" {
		return nil
	}

	if err := app.deleteObject(objectKey); err != nil {
		return err
	}
	return nil
}

func (app *S3App) MoveResources(sourceUser *userDomain.User, targetUser *userDomain.User) (userDomain.ImageUrl, userDomain.ImageUrl, error) {
	iconUrl := userDomain.ImageUrl("")
	bgImageUrl := userDomain.ImageUrl("")

	if sourceUser.IconUrl != "" {
		iconObjectKey := generateIconObjectKey(targetUser)
		err := app.moveObject(urlToObjectKey(sourceUser.IconUrl.String()), iconObjectKey)
		if err != nil {
			return "", "", err
		}

		iconUrl = userDomain.ImageUrl(objectKeyToUrl(iconObjectKey))
	}

	if sourceUser.BgImageUrl != "" {
		bgImageObjectKey := generateBgImageObjectKey(targetUser)
		err := app.moveObject(urlToObjectKey(sourceUser.BgImageUrl.String()), bgImageObjectKey)
		if err != nil {
			return "", "", err
		}

		bgImageUrl = userDomain.ImageUrl(objectKeyToUrl(bgImageObjectKey))
	}

	return iconUrl, bgImageUrl, nil
}
