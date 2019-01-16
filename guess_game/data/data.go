package data

import (
	"log"
	"fmt"

	"slice_code_test/guess_game/random"

	"github.com/jinzhu/gorm"
	// to user the db driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// UserPlayer the person playing the game
type UserPlayer struct {
	gorm.Model
	Name string `gorm:"unique"`
	Wins int
}
// Config is a customize gorm database that 
// need to be close in the main function
func Config() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./guess_game.db")
	if err != nil{
		log.Fatalln(err)
	}
	db.AutoMigrate(&UserPlayer{})
	return db
}
// createUser will add the user and win if the name is not
// in the DB
func (u *UserPlayer) createUser(s interface{}) error {
	err := Config().Create(s).Error
	if err == nil {
		return nil
	}
	return &random.CustomError{Param1: "while saving to the db =", Param2: s, Err: err}
}

// updateUser will update the win value of that user
// if he decide to play again
func (u *UserPlayer) updateUser(name string, win int) error {
	holder := UserPlayer{Name:name, Wins: win}
	Config().Where("name = ?", name).First(&u)
	err := Config().Model(&u).Update(holder).Error
	if err == nil {
		return nil
	}
	return &random.CustomError{Param1: " while updating to the db =", Param2:name, Err: err}
}

type cUDB interface {
	createUser(interface{}) error
	updateUser(string, int) error
}

// CreateUpdate is use to check if the name is in the DB
// or not then create and update it in the Database
func CreateUpdate(name string, win int, t cUDB) error {
	var us  UserPlayer
	us.Name = name
	us.Wins = win
	fmt.Println(us)
	// note it return nil because jerry is in the database
	if err := Config().Where("name = ?", us.Name).Error; err != nil{
		return t.createUser(&us)
	}
	return t.updateUser(name, win)
}
