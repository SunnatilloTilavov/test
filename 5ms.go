package main

import (
	"fmt"
	"os"
)

type Student struct {
	Name   string
	Grades []int
	Age    int
}

func main() {
	student := Student{
		Name:   "abd",
		Age:    15,
		Grades: []int{50, 60, 90, 0},
	}
	s := 0
	b := ""
	for i := 0; i < len(student.Grades); i++ {
		s += student.Grades[i]
	}
	a := s / (len(student.Grades))
	if a > 60 {
		b = "you are passed"
	} else {
		b = "you are failed"
	}

	file, err := os.Create("resault")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(b)
	if err != nil {
		fmt.Println("err", err)
		return

	}
	defer file.Close()
}
