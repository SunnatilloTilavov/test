package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	filePath := ""
	fmt.Println("file name:")
	fmt.Scan(&filePath)
	rec := 0
	fmt.Println("rec:")
	fmt.Scan(&rec)
	a := make([]string, rec)
	if checkFileExists(filePath + ".txt") {
		for i := 0; i <= 5; i++ {
			rec2 := filePath + strconv.Itoa(rand.Intn(100))
			if checkFileExists(rec2 + ".txt") {
				fmt.Println("recname", rec2)
				a = append(a, rec2)
			}
			if i == rec {
				break

			}
		}
	} else {
		file1, err := os.Create(filePath + ".txt")
		if err != nil {
			panic(err)
		}
		defer file1.Close()
	}
	filename := ""
	fmt.Println("faylga nom tanlang")
	fmt.Scan(&filename)
	file, err := os.Create(filename + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}
