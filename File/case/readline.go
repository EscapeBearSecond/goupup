package file

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 先一次性读取所有文件，然后再按行打印
func Readline1() {
	file, err := os.Open("go.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
}

// bufio 通过对io模块的封装，提供了数据的缓冲功能，能一定程度上减少大数据库读写带来的开销
// 当发起读操作时，会尝试从缓冲区读取数据，缓冲区没有数据后，才会从数据源获取
// 缓冲区大小默认为4k
func Readline2() {
	file, err := os.Open("go.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(line)
	}
}

// 通过scanner按行读取
// 单行默认大小64k
func Readline3() {
	file, err := os.Open("go.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}