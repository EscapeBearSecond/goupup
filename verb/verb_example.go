package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// pair<type:*file, value:"/dev/tty"文件描述符>
	ttf, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("Open file error, error = " + err.Error())
	}
	defer ttf.Close()
	// pair<type:, value:>
	var reader io.Reader
	// pair<type:*file, value:"/dev/tty"文件描述符>
	reader = ttf

	var writer io.Writer
	// pair<type:*file, value:"/dev/tty"文件描述符>
	writer = reader.(io.Writer)
	writer.Write([]byte("HELLO TTF\n"))
	var buffer = make([]byte, 1024)
	reader.Read(buffer)
	log.Println(string(buffer))
}
