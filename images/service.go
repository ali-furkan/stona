package images

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"github.com/ali-furkqn/stona/config"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

func Decode(src io.Reader, t string) (img image.Image, err error) {
	switch t {
	case "webp":
		{
			img, err = webp.Decode(src)
			break
		}
	default:
		{
			img, _, err = image.Decode(src)
			break
		}
	}
	return
}

func Encode(src image.Image, t string) (buf bytes.Buffer, err error) {
	switch t {
	case "jpeg":
		{
			err = jpeg.Encode(&buf, src, nil)
			break
		}
	case "png":
		{
			err = png.Encode(&buf, src)
			break
		}
	case "gif":
		{
			err = gif.Encode(&buf, src, nil)
			break
		}
	case "webp":
		{
			err = webp.Encode(&buf, src, &webp.Options{
				Lossless: false,
			})
			break
		}
	default:
		{
			err = png.Encode(&buf, src)
			break
		}
	}
	return
}

func ReSize(size uint, src io.Reader, t string) (buf bytes.Buffer, err error) {
	var (
		img    image.Image
		maxRes = config.Config().ImgMaxResolution
	)

	if size > uint(maxRes) {
		return buf, errors.New("Size not allowed. Size must be smaller than " + fmt.Sprint(maxRes))
	}

	img, err = Decode(src, t)

	if err != nil {
		return buf, err
	}

	imgSize := img.Bounds().Size()
	ratio := float64(imgSize.X) / float64(imgSize.Y)

	newImage := resize.Resize(uint(ratio*float64(size)), uint(size), img, 2)

	buf, err = Encode(newImage, t)

	return buf, err
}
