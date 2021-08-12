package main

import (

	"fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func(mReader MyReader) Read(slice []byte)( n int ,err error){
	slice = slice[:cap(slice)];
	fmt.Println(len(slice))
	for i:=range slice{
	slice[i] = 'A'
	}
	return cap(slice) ,err
}

func main() {
	reader.Validate(MyReader{})
}
