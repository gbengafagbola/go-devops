package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	Contents string
	Position int
}


func (m MySlowReader) Read(p []byte) (n int, err error){
	return 0, io.EOF 
} 

func main() {
	mySlowReaderInstance := MySlowReader{
		Contents: "hello",
	}  

	body, err := io.ReadAll(mySlowReaderInstance)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("output: %s \n", body)
}