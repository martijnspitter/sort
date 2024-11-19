package reader

import (
	"io"
	"os"
)

type Reader interface {
	Read(path string) (io.Reader, error)
}

type FileReader struct {
}

func NewReader() *FileReader {
	return &FileReader{}
}

func (f *FileReader) Read(path string) (io.Reader, error) {
	if path == "" {
		return os.Stdin, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}
