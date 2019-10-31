package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	var result byte
	switch {
	case b >= 'A' && b <= 'Z':
		result = (b-'A'+13)%26 + 'A'
	case b >= 'a' && b <= 'z':
		result = (b-'a'+13)%26 + 'a'
	default:
		result = b
	}
	return result
}

func (rot *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}


