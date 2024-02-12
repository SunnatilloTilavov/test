package main

import (
 "fmt"
)

type People struct {
 Name  string `json:"name"`
 Sex   string `json:"sex"`
 Job string `json:"class"`

}

func main() {
	var name, sex, job,seex string
var son int
 people := make([]People, 0)
 fmt.Print("sonni:")
 fmt.Scan(&son)

 for i := 1; i <= son; i++ {
 fmt.Println(i,"-odam")
  fmt.Print("Name: ")
  fmt.Scan(&name)

  fmt.Print("Sex: ")
  fmt.Scan(&sex)

  fmt.Print("Job: ")
  fmt.Scan(&job)

  people = append(people, People{Name: name, Sex: sex, Job: job})
 }
 fmt.Println("Jinsini  male,female,all")
 fmt.Scan(&seex)
for _,peop:=range people{
	if peop.Sex==seex{
		fmt.Printf("Name:%v; sex:%v; job:%v \n",peop.Name,peop.Sex,peop.Job)

	}else if seex=="all"{
		fmt.Printf("name:%v; sex:%v; job:%v \n",peop.Name,peop.Sex,peop.Job)

	}else{
		fmt.Print("jinsini noto'gri kirittingiz")
	}
}

}