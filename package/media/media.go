package media

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/nfnt/resize"
)

type Media struct{}

func New() IMedia {
	return &Media{}
}

func (m *Media) ResizeImage(width int64, height int64, r io.Reader) (*bytes.Buffer, error) {
	img, ext, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	val := resize.Resize(uint(width), 0, img, resize.Lanczos3)

	w := new(bytes.Buffer)

	if ext == "jpeg" {
		if err := jpeg.Encode(w, val, nil); err != nil {
			return nil, err
		}
	}
	if ext == "png" {
		if err := png.Encode(w, val); err != nil {
			return nil, err
		}
	}

	return w, nil
}

func (m *Media) GetImageSize(r io.Reader) (int, int, error) {
	cfg, _, err := image.DecodeConfig(r)
	if err != nil {
		return 0, 0, err
	}

	return cfg.Width, cfg.Height, nil
}
