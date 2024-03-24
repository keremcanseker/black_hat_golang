package main

import (
	"fmt"
	"io"
	"log"

	"os"
)

type FooReader struct{}
func (fooReader *FooReader) Read(p []byte) (n int, err error){
	fmt.Print("in> ")
	return os.Stdin.Read(p)
}

type FooWriter struct{}
func (fooWriter *FooWriter) Write(p []byte) (n int, err error){
	fmt.Print("out> ")
	return os.Stdout.Write(p)
}


func Copy(dst io.Writer, src io.Reader) (written int64, err error){
	return io.Copy(dst, src)
}

func main(){
	var (
		reader FooReader
		writer FooWriter
	)
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
		}
}