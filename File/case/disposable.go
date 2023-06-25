package file

import (
	"log"
	"os"
	"path"
)
// 一次性读取所有文件内容并写入
func ReadWriteFile() {
	list := getFiles(sourceDir)

	for _, f := range list {
		bytes, err := os.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}
		_, name := path.Split(f)
		destName := destDir + "disposable/" + name
		err = os.WriteFile(destName, bytes, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}