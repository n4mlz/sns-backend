package images

import (
	"bytes"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

const IMAGE_SIZE = 128

func resizeImage(img image.Image, width, height int) image.Image {
	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	draw.BiLinear.Scale(newImage, newImage.Bounds(), img, img.Bounds(), draw.Over, nil)

	return newImage
}

func fotmatImage(file *os.File) ([]byte, error) {
	img, _, err := image.Decode(file)
	if err != nil || img == nil {
		return nil, err
	}

	resizedImg := resizeImage(img, IMAGE_SIZE, IMAGE_SIZE)

	var buf bytes.Buffer
	if err := png.Encode(&buf, resizedImg); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
