package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

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
	//dairyFile, err := os.Create("Dairy_file")
	dairyFile, err := os.OpenFile("Dairy_file", os.O_APPEND|os.O_WRONLY, 0666)
	defer dairyFile.Close()
	check(err)
	for _, hr := range holdRecord {
		json.NewEncoder(dairyFile).Encode(&hr)
	}
}

func readDataJSON() Record {
	var p Record
	file, err := os.Open("Dairy_file")
	check(err)
	defer file.Close()
	json.NewDecoder(file).Decode(&p)

	return p
}

func filter(s Record, f func(Record) bool) Record {

	return s

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

func getUserIput() []Record {
	var r Record
	var sliceRecord []Record
	fmt.Println("Enter name:")
	r.Name = giveText()
	fmt.Println("Enter secred key: ")
	r.SecredKey = giveText()
	fmt.Println("Enter Diary:")
	r.Diary = giveText()
	sliceRecord = append(sliceRecord, r)
	return sliceRecord
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
