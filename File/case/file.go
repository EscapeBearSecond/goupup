package file

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const sourceDir = "File/source/"
const destDir = "File/dest/"

func getFiles(dir string) []string {
	fs, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	list := make([]string, 0)
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		fullName := strings.Trim(dir, "/") + "/" + f.Name()
		list = append(list, fullName)
	}
	fmt.Println(list)
	return list
}