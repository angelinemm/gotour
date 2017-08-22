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

var upperASCII = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lowerASCII = []byte("abcdefghijklmnopqrstuvwxyz")

func rot13(b byte) byte {
    newChar := b
    foundUpper := bytes.IndexByte(upperASCII, b)
	foundLower := bytes.IndexByte(lowerASCII, b)
	if foundUpper != -1 {
	    newChar = upperASCII[(foundUpper + 13) % 26]
	}
	if foundLower != -1 {
	    newChar = lowerASCII[(foundLower + 13) % 26]
	}
	return newChar
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
