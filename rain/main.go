package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// problem: Create a string array containing a representation of the 52 cards in a standard deck, and then shuffles the 52 cards into a second array.
//Start with two string arrays of the cards and suits:
//Values: 2, 3, 4, 5, 6, 7, 8, 9, 10, J, Q, K, A
//Suits: Clubs, Diamonds, Hearts, Spades
//Image of all 52 cards https://bit.ly/3aC9Kep
//Sample input: [‘2’,’3’,...’K’, ‘A’] [‘Clubs’, ...’Diamonds’]
//Sample output: [‘2Clubs’, ‘2Hearts’,...]

// Select any general data model you might be familiar
//with (e.g. NBA teams, social media platforms, cryptocurrencies, etc.)
//Setup a basic RESTful API web service that has an endpoint that can either accept a request to create a new record, or a request where it will return all the items in a collection.
//Pose the question of if we were to use a database instead, how would they handle it?

type NBATeams struct {
	Name    string `json:"name"`
	Players string `json:"players"`
}

type DB map[string]*NBATeams

func main() {
	router := gin.Default()
	db := make(DB)
	var teams NBATeams
	router.POST("/add", func(context *gin.Context) {
		if err := context.ShouldBindJSON(&teams); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "error add teams"})
			return
		}
		db[teams.Name] = &teams
		context.JSON(http.StatusOK, gin.H{"success": teams})
	})

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"success": db})
	})

	if err := router.Run(":3002"); err != nil {
		log.Fatalln("error starting up server")
	}
}
