package main

import (
	"fmt"
	//"strings"
	//"io/ioutil"
	"encoding/json"
	"log"
	"os"
)

type person struct {
	Name   string
	Email  string
	Age    int `json:"old"`
	Gender string
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func main() {
	var p person
	p.Age = 29
	p.Name = "jerry"
	p.Gender = "male"
	p.Email = "jerry@gmail.com"
	f, err := os.Create("my_stream_of_json")
	check(err)
	// json.NewEncoder(os.Stdout).Encode(p)
	defer f.Close()
	json.NewEncoder(f).Encode(&p)

	// i have create a file now is time to read it as a go type
	var p1 streamPerson

	// this works i use ioutil and string.NewReader(string())

	/* fileByte, err := ioutil.ReadFile("my_stream_of_json")
	check(err)
	jsonObj := json.NewDecoder(strings.NewReader(string(fileByte)))
	jsonObj.Decode(&p1) */

	// i want to try using os.Open to read the file

	fileByte, err := os.Open("my_stream_of_json")
	check(err)
	json.NewDecoder(fileByte).Decode(&p1)
	fmt.Println(p1.Age, p1.Name, p1.Email, p1.Gender)
	fileByte.Close()

}

type streamPerson struct {
	Name   string
	Email  string
	Age    int `json:"old"`
	Gender string
}
