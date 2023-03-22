package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
func (r13 rot13Reader) Read(p []byte) (int,error){
	n, err := r13.r.Read(p)
	for i := 0; i < n; i++ {
		if (p[i] >= 'A' && p[i] <= 'M') || (p[i] >= 'a' && p[i] <= 'm') {
			p[i] += 13
		} else if (p[i] >= 'N' && p[i] <= 'Z') || (p[i] >= 'n' && p[i] <= 'z') {
			p[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
