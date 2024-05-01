package media

import (
	"bytes"
	"io"
)


type IMedia interface {
	GetImageSize(r io.Reader) (int, int, error)
	ResizeImage(width int64, height int64, r io.Reader) (*bytes.Buffer, error)
}
