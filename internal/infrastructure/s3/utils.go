package s3

import (
	"bytes"
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

func resizeImage(img image.Image, width, height int) image.Image {
	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.BiLinear.Scale(newImage, newImage.Bounds(), img, img.Bounds(), draw.Over, nil)

	return newImage
}

func fotmatImageForIcon(file io.Reader) ([]byte, error) {
	img, _, err := image.Decode(file)
	if err != nil || img == nil {
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
	img, _, err := image.Decode(file)
	if err != nil || img == nil {
		return nil, err
	}

	resizedImg := resizeImage(img, BG_IMAGE_WIDTH, BG_IMAGE_HEIGHT)

	var buf bytes.Buffer
	if err := png.Encode(&buf, resizedImg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
