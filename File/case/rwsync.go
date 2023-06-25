package file

import (
	"io"
	"log"
	"os"
	"path"
)

func RwSync() {
	list := getFiles(sourceDir)
	for _, f := range list {
		_, name := path.Split(f)
		destName := destDir + "rwsync/" + name
		writeWhileReading(f, destName)
	}
}

// 边读边写，适合大文件的读写
func writeWhileReading(srcName, destName string) {
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	var buff = make([]byte, 1024)
	dest, err := os.OpenFile(destName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for {
		n, err := src.Read(buff)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n <= 0 {
			break
		}
		_, err = dest.Write(buff[:n])
		if err != nil {
			log.Fatal(err)
		}
	}
}
