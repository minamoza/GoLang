package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (int,error) {
	n, err := r.r.Read(b)
	
	for i:= 0; i < len(b); i++{
		if b[i] >= 'a' && b[i] <= 'z' {
			if b[i] >= 'm' {
				b[i] -= 13
			} else {
				b[i] += 13
			}
		} else if b[i] >= 'A' && b[i] <= 'Z' {
			if b[i] >= 'M' {
				b[i] -= 13
			} else {
				b[i] += 13
			}
		}
	}
    return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
