package s3

import (
	"bytes"
	"errors"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/draw"
)

const (
	ICON_IMAGE_WIDTH  = 256
	ICON_IMAGE_HEIGHT = 256
)

const (
	BG_IMAGE_WIDTH  = 960
	BG_IMAGE_HEIGHT = 320
)

const (
	maxImageBytes   = 10 << 20
	maxImagePixels  = 64 * 1024 * 1024
	maxImageSideLen = 8192
)

func resizeImage(img image.Image, width, height int) image.Image {
	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.BiLinear.Scale(newImage, newImage.Bounds(), img, img.Bounds(), draw.Over, nil)

	return newImage
}

func fotmatImageForIcon(file io.Reader) ([]byte, error) {
	img, err := decodeSafeImage(file)
	if err != nil {
		return nil, err
	}

	resizedImg := resizeImage(img, ICON_IMAGE_WIDTH, ICON_IMAGE_HEIGHT)

	var buf bytes.Buffer
	if err := png.Encode(&buf, resizedImg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func fotmatImageForBgImage(file io.Reader) ([]byte, error) {
	img, err := decodeSafeImage(file)
	if err != nil {
		return nil, err
	}

	resizedImg := resizeImage(img, BG_IMAGE_WIDTH, BG_IMAGE_HEIGHT)

	var buf bytes.Buffer
	if err := png.Encode(&buf, resizedImg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func decodeSafeImage(file io.Reader) (image.Image, error) {
	data, err := io.ReadAll(io.LimitReader(file, maxImageBytes+1))
	if err != nil {
		return nil, err
	}
	if len(data) > maxImageBytes {
		return nil, errors.New("image is too large")
	}

	config, format, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil {
		return nil, errors.New("invalid image")
	}
	if format != "png" && format != "jpeg" {
		return nil, errors.New("unsupported image format")
	}
	if config.Width < 1 || config.Height < 1 || config.Width > maxImageSideLen || config.Height > maxImageSideLen || config.Width > maxImagePixels/config.Height {
		return nil, errors.New("image dimensions are too large")
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, errors.New("invalid image")
	}
	return img, nil
}
