package main

import (
	"fmt"
)

func main() {

	var bldr = newNotificationBuilder()

	bldr.SetTitle("Hello!!!")
	bldr.SetPriority(10)

	notification, err := bldr.Build()
	if err != nil {
		fmt.Println("You have an error", err)
	}

	fmt.Printf("Notification: %+v", notification)
}
