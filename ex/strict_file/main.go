package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)
const fileName = "d.json"
// Record of my files
type Record struct {
	Name      string
	SecredKey string
	Diary     string
}

func check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func giveText() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	check(err)
	return text
}

func writeDataJSON() {
	holdRecord := getUserIput()
	var hR []Record
	hR = append(hR, holdRecord)
	//os.Create(fileName)
	dairyFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0666)
	defer dairyFile.Close()
	check(err)
	json.NewEncoder(dairyFile).Encode(&hR)
}

func readDataJSON() []Record {
	var p []Record
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	json.NewDecoder(file).Decode(&p)
	log.Println(p)
	return p
}

func filter(s []Record, f func(Record) bool) []Record {
	var userRecord []Record
	for _, r := range s {
		if f(r) {
			userRecord = append(userRecord, r)
		}
	}
	return userRecord

}

func choiceOption(one int) {
	switch one {
	case 1:
		writeDataJSON()
		fmt.Println("Success")
		ms()
		choiceOption(start())
	case 2:
		fmt.Print("Enter your secret Key: ")
		secretKey := giveText()
		// filter call
		record := filter(readDataJSON(), func(r Record) bool {
			if r.SecredKey == fmt.Sprintf("%s%s", secretKey, "\n") {
				return true
			}
			return false
		})
		fmt.Println(record)
		ms()
		choiceOption(start())
	case 3:
		fmt.Println("Good bye")
		os.Exit(0)

	default:
		fmt.Println("Not a number")
		os.Exit(0)
	}

}

func getUserIput() Record {
	var r Record
	fmt.Println("Enter name:")
	r.Name = giveText()
	fmt.Println("Enter secred key: ")
	r.SecredKey = giveText()
	fmt.Println("Enter Diary:")
	r.Diary = giveText()
	return r
}

func ms() {
	fmt.Println("Hello welcome to my Diary Book\n\t\tEnter\t1:to input date\n\t\t\t2:To read your input\n\t\t\t3:to close.")
}

func start() int {
	var start int
	fmt.Print("Enter your option: ")
	fmt.Scan(&start)
	return start
}

func main() {
	ms()
	choiceOption(start())
}
