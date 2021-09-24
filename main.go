// Check Directory and Files includes image(*.bmp | .jpg | .png) and json
package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// Find recursive bmp | jpg and json given path
// Print Result And Save CSV
func PrintDirInfo(path string, w *csv.Writer) {
	if files, err := ioutil.ReadDir(path); err != nil {
		log.Fatal(err)
	} else {
		var imgs int
		var labels int
		for _, file := range files {
			if file.IsDir() {
				PrintDirInfo(path+"\\"+file.Name(), w)
			} else {
				fileName := file.Name()
				if CheckImageFile(fileName) {
					imgs += 1
				} else if strings.HasSuffix(fileName, "json") {
					labels += 1
				}
			}
		}
		w.Write([]string{path, fmt.Sprint(len(files)), fmt.Sprint(imgs), fmt.Sprint(labels)})
		fmt.Printf("%s, total: %d, imgs: %d, labels: %d\n", path, len(files), imgs, labels)
	}
}

func CheckImageFile(fileName string) bool {
	pat := regexp.MustCompile(`(bmp|BMP|jpg|JPG|png|PNG)$`)
	return pat.MatchString(fileName)
}

func main() {
	// 경로를 입력받는다
	if len(os.Args) < 2 {
		return
	}
	dirName := os.Args[1]

	// CSV 파일
	if file, err := os.Create("output.csv"); err != nil {
		log.Fatal(err)
	} else {
		csvWriter := csv.NewWriter(file)
		PrintDirInfo(dirName, csvWriter)
		csvWriter.Flush()
	}
}
