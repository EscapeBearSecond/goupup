package main

import (
	"fmt"
	"io"
	"os"
	"syscall"
)

func main() {
	// 打开文件
	//openFile()
	// 写文件
	//writeFile()
	// 读文件
	//readFile()
	copyFile()
}
func copyFile() {
	srcFile, err := os.OpenFile("./hello.txt", syscall.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	dstFile, err := os.OpenFile("./hello2.txt", syscall.O_CREAT|syscall.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, err := srcFile.ReadAt(buf[:], 10)
		if err == io.EOF {
			fmt.Printf("读取完毕\n")
			dstFile.Write(buf[:n])
			break
		}
		if err != nil {
			fmt.Printf("read srcFile failed, err: %v\n", err)
			return
		}
		dstFile.Write(buf[:n])
	}
	srcFile.Close()
	dstFile.Close()
}
func readFile() {
	file, err := os.Open("./hello.txt")
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
		return
	}
	defer file.Close()
	var buf [128]byte
	var content []byte
	for {
		//n, err := file.ReadAt(buf[:], 10) // 本次调用就直接返回io.EOF
		n, err := file.Read(buf[:]) // 在下一次调用返回io.EOF
		fmt.Println(n)
		if err == io.EOF {
			fmt.Println("read file over....")
			//content = append(content, buf[:n]...)
			break
		}
		if err != nil {
			fmt.Printf("read file failed, err: %v\n", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}
func openFile() {
	file, err := os.Open("./example-10.go")
	if err != nil {
		fmt.Printf("open file failed err: %v\n", err)
		return
	}
	file.Close()
}
func writeFile() {
	file, err := os.OpenFile("./hello.txt", syscall.O_CREAT|syscall.O_APPEND|syscall.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("open file failed err: %v\n", err)
		return
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
		file.WriteString("ab\n")
		file.Write([]byte("cd\n"))
	}
}
