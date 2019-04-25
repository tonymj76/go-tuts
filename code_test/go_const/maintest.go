package main

import (
	"path/filepath"
	"os"
	"crypto/rand"
	"fmt"
	"encoding/hex"
)

func main() {
	len := 20
	buff := make([]byte, len)
	// here is where all the magic takes place
	// here i use crypto instead of math to gen run output then before
	rand.Read(buff)
	fmt.Println(buff)
	str := hex.EncodeToString(buff)
	fmt.Println(str)
	fmt.Println()

	holder, _ := os.Getwd()
	holder = filepath.Join(holder, "/tony/mj/", "somethin.jpg")
	fmt.Println(holder)
	os.MkdirAll(holder, 0766)
}