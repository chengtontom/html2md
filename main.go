package main

import (
	"flag"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func convertHtml2Md(path string,f os.FileInfo) error {

	if f.Name()[len(f.Name())-5:] != ".html" {
		return nil
	}

	mdFileName := f.Name()[:len(f.Name())-5] + ".md"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	print(string(data))

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(string(data))
	if err != nil {
		log.Fatal(err)
	}

	mdFile, err := os.OpenFile(
		mdFileName,
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写字节到文件中

	bytes , err := mdFile.Write([]byte(markdown))
	if err != nil {
		log.Fatal(err)
	}

	println("write %d byes", bytes)

	return err
}

func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(path)

		convertHtml2Md(path, f)

		return nil
	})

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func main() {

	flag.Parse()
	root := "."//flag.Arg(0)
	getFilelist(root)

}
