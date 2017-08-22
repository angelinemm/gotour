package main

import (
	"io"
	"os"
	"strings"
	"bytes"
)

type rot13Reader struct {
	r io.Reader
}

var upper_ascii = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lower_ascii = []byte("abcdefghijklmnopqrstuvwxyz")

func rot13(b byte) byte {
    new_char := b
    found_upper := bytes.IndexByte(upper_ascii, b)
	found_lower := bytes.IndexByte(lower_ascii, b)
	if found_upper != -1 {
	    new_char = upper_ascii[(found_upper + 13) % 26]
	}
	if found_lower != -1 {
	    new_char = lower_ascii[(found_lower + 13) % 26]
	}
	return new_char
}


func (r rot13Reader) Read(stream []byte) (n int, err error){
    n, err = r.r.Read(stream)
    for i := 0; i<n; i++ {
	    stream[i] = rot13(stream[i])
	}
	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
