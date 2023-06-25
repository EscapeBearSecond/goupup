package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"syscall"
)

func main() {
	//writeFile()
	readFile()

}

func writeFile() {
	file, _ := os.OpenFile("./hello.txt", syscall.O_RDWR|syscall.O_CREAT, 0666)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n")
	}
	writer.Flush()
}
func readFile() {
	file, _ := os.OpenFile("./hello.txt", syscall.O_RDONLY, 0666)
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}
}
