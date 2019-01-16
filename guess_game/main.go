package main

import (
	"fmt"
	"log"
	"slice_code_test/guess_game/data"
	"slice_code_test/guess_game/random"
)

func main() {
	var (
		name string
		err  error
		win  int
		try  int
		db   data.UserPlayer
	)

	randomNumber := random.GenRanNum(100)
	fmt.Println(randomNumber)

	name = random.GetUserName("Name: ")
	for {
		guessedNumber, _ := random.Input("enter a number: ")
		if guessedNumber > 100 {
			fmt.Println(guessedNumber, "is above 100 you should guess between 1 and 99")
		} else {
			if guessedNumber == randomNumber {
				fmt.Printf("%v won the number is %v\n", name, guessedNumber)
				win = 4
				err = data.CreateUpdate(name, win, &db)
				break
			} else {
				if randomNumber < guessedNumber {
					fmt.Printf("%v is higher than the number guessed\n", guessedNumber)
					try++
				} else {
					fmt.Printf("%v is lower than the number guessed\n", guessedNumber)
					try++
				}
			}
		}

	}
	if err != nil{
		log.Fatal(err)
	}
	// here i close all connection to database
	defer data.Config().Close()
}
