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
		if b[i] >= 'A' && b[i] <= 'Z'{
			b[i] = byte((b[i]-'A'+13)%26 + 'A')
		} else if b[i] >= 'a' && b[i] <= 'z'{
			b[i] = byte((b[i]-'a'+13)%26 + 'a')
		}
	}
    return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
