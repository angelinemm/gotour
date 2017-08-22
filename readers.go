package main

import "golang.org/x/tour/reader"

// MyReader is a reader
type MyReader struct{}

// Read generates an endless stream of 'A' characters
func (MyReader) Read(stream []byte) (n int, err error){
    stream[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
