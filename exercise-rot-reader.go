package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	l, err := r.r.Read(b)
	for i := 0; i < l; i++ {
		if 'A' <= b[i] && b[i] <= 'M' || 'a' <= b[i] && b[i] <= 'm' {
			b[i] += 13
		} else if 'N' <= b[i] && b[i] <= 'Z' || 'n' <= b[i] && b[i] <= 'z' {
			b[i] -= 13
		}
	}
	return l, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
