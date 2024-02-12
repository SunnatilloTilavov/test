package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	// "path/filepath"
	// "strconv"
)

func main() {

	filePath := ""
	fmt.Println("file name:")
	fmt.Scan(&filePath)
	rec := 0
	fmt.Println("rec:")
	fmt.Scan(&rec)
	a := make([]int, rec)
	if !checkFileExists(filePath + ".text") {
		for i := 0;i<=5; i++ {
			rec2 := filePath + strconv.Itoa(rand.Intn(100)) + ".txt"
			if checkFileExists(rec2) {
				a=append(a,rec)
				fmt.Println("recname", rec2)
			
			}else{
				_, err := os.Create( rec2)
				if err != nil {
					panic(err)
				}

			} 
			if len(a) == rec {
				break

			}
		}
	}else {
		_, err := os.Create(filePath + ".txt")
		if err != nil {
			panic(err)
		}
	}
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}
