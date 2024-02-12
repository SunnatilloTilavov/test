
package main
 
import (
    "fmt"
    "os"
    "bufio"
    "log"
	"strconv"
	"strings"
)
 
func main() {
	a:="numb.txt"
	b:="words.txt"
		file1, err := os.Create(a)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file1.Close() 

		file2, err := os.Create(b)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file2.Close() 

 
    file, err := os.Open("example.txt")
    if err != nil {
        log.Fatal(err)
    }
 
    Scanner := bufio.NewScanner(file)
    Scanner.Split(bufio.ScanWords)
 var s int
	for Scanner.Scan() {
		word := Scanner.Text()

		if isInteger(word) {
			num, err := strconv.Atoi(word)
			if err != nil {
				fmt.Println(err)
			} else {
				s+=num
				_, err := file1.WriteString(strconv.Itoa(num) + "\n")
				if err != nil {
				  fmt.Println("err", err)
				  return
				}
			}
		} else if  strings.Contains(word, "A")||strings.Contains(word, "O")||strings.Contains(word, "a")||strings.Contains(word, "o")    {
			_, err := file2.WriteString(word + "\n")
			if err != nil {
			  fmt.Println("error write word:", err)
			  return
		}

	}
    if err := Scanner.Err(); err != nil {
        log.Fatal(err)
    }

}
 _, err = file1.WriteString(strconv.Itoa(s) + "\n")
 if err != nil {
fmt.Println("err", err)
  return

}
defer file.Close()
defer file1.Close()
}



func isInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}