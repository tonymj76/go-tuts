package main

import (
	"fmt"
	"log"
	"encoding/json"
)

type identification struct {
	Name  string
	Phone int
	Email string
}
func (n *identification) String() string{
	return fmt.Sprintf()
}
func main() {
	jsonText := []byte(`[
        {"Name": "ID1", "Phone": 0, "Email": "email@email.com"}
	 ]`)
	 
	var idents []identification
	if err := json.Unmarshal(jsonText, &idents); err != nil{
		log.Println(err)
	}
	fmt.Println(idents)

}