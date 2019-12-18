package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

// 画像かどうかを判定する。
// 拡張子で判定するのではなく、画像自体を判定する。
// 対応画像: gif、jpeg、png、bmp、tiff
func main() {
	var arg string
	if len(os.Args) != 2 {
		fmt.Println("The number of arguments specified is incorrect.")
		os.Exit(1)
	} else {
		arg = os.Args[1]
	}

	paths := dirwalk(arg)

	for _, path := range paths {
		fmt.Println(path)
		f, _ := os.Open(path)
		defer f.Close()

		_, format, err := image.DecodeConfig(f)

		if err != nil {
			fmt.Println(err)
			//return
		}

		fmt.Println(format)
		fmt.Println("")
	}
}
