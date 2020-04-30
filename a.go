package main

import (
	"golang.org/x/tour/reader"
	"io"
)

type MyReader struct{}

// Read ...
func (m MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return 0, io.ErrShortBuffer
	}
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
