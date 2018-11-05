package main

import (
	"encoding/json"
	"fmt"
	"os"
	"bufio"
	"log"
)
// Record of my files
type Record struct {
	Name 			string
	SecredKey 	string
	Diary 		string
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

func init() {
	fmt.Println("Hello welcome to my Diary Book\n\t\tEnter\t1:to input date\n\t\t\t2:To read your input\n\t\t\t3:to close.")
}

func writeDataJSON(start int) {
	holdRecord := inputData(start)
	//dairyFile, err := os.Create("Dairy_file")
	dairyFile, err := os.OpenFile("Dairy_file", os.O_APPEND|os.O_WRONLY, 0666)
	check(err)
	for _, hr := range holdRecord {
		json.NewEncoder(dairyFile).Encode(&hr)
	}
}

func filter(s []Record, f func(Record) bool) []Record{
	var r []Record
	for _, k := range s {
		if f(k) {
			r = append(r, k)
		}
	}
	return r

}

func readDataJSON(start int) []Record{
	var p []Record
	file, err := os.Open("Dairy_file")
	check(err)
	defer file.Close()
	json.NewDecoder(file).Decode(&p)
	return p
}

func inputData(one int) []Record{
	var r Record
	var sliceRecord []Record
	if one==1 {
		fmt.Println("Enter name:")
		r.Name = giveText()
		fmt.Println("Enter secred key: ")
		r.SecredKey = giveText()
		fmt.Println("Enter Diary:")
		r.Diary = giveText()
		sliceRecord = append(sliceRecord, r)
	}
	return sliceRecord
}

func main() {
	// Show the menu
	// Recieve input
	// Switch
	
	var start int
	fmt.Scan(&start)
	writeDataJSON(start)
	secretKey := giveText()

	filter(readDataJSON(1), func (r Record) bool  {
		if r.SecredKey == secretKey {
			return true
		}
		return false
	})
}