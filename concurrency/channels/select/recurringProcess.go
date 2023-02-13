package main

import (
	"fmt"
	gofakeit "github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"time"
)

type Notification struct {
	UserId  int
	Content string
}

type pendingNotifications map[int][]*Notification

func sendUserBatchNotifications(userId int, notifications []*Notification) {
	fmt.Printf("Sending email to user with userId %d for pending notifications %v\n", userId, notifications)
}

func handlePendingUsersNotifications(pn pendingNotifications, handler func(userId int, notifications []*Notification)) {
	for userId, notification := range pn {
		handler(userId, notification)
		delete(pn, userId)
	}
}

func collectNewUsersNotifictions(pn pendingNotifications) {
	randomNotifications := getRandomNotifications()
	if len(randomNotifications) > 0 {
		pn[randomNotifications[0].UserId] = randomNotifications
	}
}

func getRandomNotifications() (notifications []*Notification) {
	rand.Seed(time.Now().UnixNano())
	userId := rand.Intn(100-10+1) + 10
	numOfNotifications := rand.Intn(5-0+1) + 0
	fmt.Printf("numOfNotifications %v\n numberOfUserId %d\n", numOfNotifications, userId)
	for i := 0; i < numOfNotifications; i++ {
		notifications = append(notifications, &Notification{Content: gofakeit.Paragraph(1, 2, 10, " "), UserId: userId})
	}
	return
}

func main() {

}
