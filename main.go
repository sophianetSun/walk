// Check Directory and Files includes image(*.bmp | .jpg | .png) and json
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func PrintDirInfo(path string) {
	if files, err := ioutil.ReadDir(path); err != nil {
		log.Fatal(err)
	} else {
		var imgs int
		var labels int
		for _, file := range files {
			if file.IsDir() {
				PrintDirInfo(path + "\\" + file.Name())
			} else {
				fileName := file.Name()
				if strings.HasSuffix(fileName, "jpg") || strings.HasSuffix(fileName, "bmp") {
					imgs += 1
				} else if strings.HasSuffix(fileName, "json") {
					labels += 1
				}
			}
		}
		fmt.Printf("%s, total: %d, imgs: %d, labels: %d\n", path, len(files), imgs, labels)
	}
}

func main() {
	// 경로를 입력받는다

	if len(os.Args) < 2 {
		return
	}
	dirName := os.Args[1]
	PrintDirInfo(dirName)
	// 폴더 내에 bmp 개수를 찾는다
	// 폴더 내에 json 개수를 찾는다
	// 하위 폴더를 탐색한다
	// 결과를 보여준다
	// 폴더 | 이미지 수 | 레이블 수

}
