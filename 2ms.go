package main
import (
	"fmt"
    "os"
	"encoding/json"
	"io/ioutil"
	"sort"
    // "bufio"
    // "log"
	// "strconv"
	// "strings"
)


type Info struct {
	Firstname string `json:"first_name"`
	Secondname string `json:"second_name"`
	Salary int `json:"salary"`
	Experience int `json:"experience"`
	Age int `json:"age"`



}
func main(){
	a:="orderByAge.txt"
	b:="topSalaryEmployees.txt"
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
		fileName := "employees.json"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	var info []Info
	err = json.Unmarshal(content, &info)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	sort.Slice(info, func(i, j int) bool {
		return info[i].Age <info[j].Age
	})
	offSet1:=0
	offSet:=0

	for _, m := range info {
		text:=fmt.Sprintf("firstname %v , secondname %v , salary: %v,expp: %v , age:%v \n",m.Firstname,m.Secondname,m.Salary,m.Experience,m.Age)
		written,err:=file1.WriteAt([]byte(text),int64(offSet1))

		if err!=nil{
			panic(err)
		}
		offSet1+=written
		
	}


//oylik boyicha top3

sort.Slice(info, func(i, j int) bool {
	return info[i].Salary > info[j].Salary
})

// for _, m := range info {
// 	text:=fmt.Sprintf("firstname %v , secondname %v , salary: %v,expp: %v , age:%v \n",m.Firstname,m.Secondname,m.Salary,m.Experience,m.Age)
// 	written,err:=file2.WriteAt([]byte(text),int64(offSet))

// 	if err!=nil{
// 		panic(err)
// 	}
// 	offSet+=written
// }
for i:=0;i<=2;i++ {

	text:=fmt.Sprintf("firstname %v , secondname %v , salary: %v,expp: %v , age:%v \n",info[i].Firstname,info[i].Secondname,info[i].Salary,info[i].Experience,info[i].Age)
	written,err:=file2.WriteAt([]byte(text),int64(offSet))

	if err!=nil{
		panic(err)
	}
	offSet+=written
}
		
}
