package random

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
	"time"
	"math/rand"
)

// GenRanNum is used to generate random numbers
func GenRanNum(nums int) (num int) {
	varyingSequence := rand.NewSource(time.Now().UnixNano())
	newVS := rand.New(varyingSequence)
	num = newVS.Intn(nums)
	return
}

// Input takes in a message and the user input and return an int
func Input(msg string) (int, error){
	fmt.Print(msg)
	var userInput string
	fmt.Scan(&userInput)
	return strconv.Atoi(userInput)
}

// GetUserName the user name
func GetUserName(msg string) string {
	fmt.Print(msg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
	
}