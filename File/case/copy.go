package file

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func CopyDirToDir() {
	list := getFiles(sourceDir)
	for _, f := range list {
		_, name := path.Split(f)
		fmt.Println("fileName is", name)
		destFileName := destDir + "copy/" + name
		fmt.Println("destName is", destFileName)
		CopyFile(f, destFileName)
	}
}
// 拷贝文件，而不是拷贝目录
func CopyFile(srcName, destName string) (int64, error) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	dest, err := os.OpenFile(destName, os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()
	return io.Copy(dest, src)
}