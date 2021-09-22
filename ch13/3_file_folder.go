package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(fileName string) {
	// bs, err := ioutil.ReadFile("1_strings.go")
	bs, err := os.ReadFile(fileName) // from go 1.16
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("test")
}

func ReadDir(dirPath string) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return
	}
	defer dir.Close()

	// fileInfos, err := dir.Readdir(-1)
	// if err != nil {
	// 	return
	// }
	// for _, fi := range fileInfos {
	// 	fmt.Println(fi.Name())
	// }

	// from go 1.16
	dirEntries, err := dir.ReadDir(-1)
	if err != nil {
		return
	}
	for _, entry := range dirEntries {
		fmt.Println(entry.Name())
	}
}

func ReadDirRecursion(dirPath string) {
	// filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
	// 	fmt.Println(path)
	// 	return nil
	// })

	// from go 1.16
	filepath.WalkDir(dirPath, func(path string, dirEntry os.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	})
}

func main() {
	ReadFile("1_strings.go")
	CreateFile("test.txt")
	//ReadDir(".")
	ReadDirRecursion(".")
}
