package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Reads the stream and rot13 encodes all alphabetic elements.
func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

// Applies the rot13 cipher to a single byte if it is alphabetical,
// if not, no transformation is applied.
func rot13(b byte) byte {
	var letter byte
	switch {
	case b >= 'a' && b <= 'z':
		letter = b - 'a'
		return ((letter + 13) % 26) + 'a'
	case b >= 'A' && b <= 'Z':
		letter = b - 'A'
		return ((letter + 13) % 26) + 'A'
	default:
		return b
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
