package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {

	//hourOfDay := time.Now().Hour()
	greeting, err := getGreeting(4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Current Time : ", time.Now().Local())
	fmt.Println(greeting)
}

func getGreeting(hour int) (string, error) {

	var message string

	if hour < 7 {
		err := errors.New("Too early in the morning!")
		return message, err
	}
	if hour < 12 {
		message = "Good Morning!"
	} else if hour < 18 {
		message = "Good Afternoon!"
	} else if hour < 21 {
		message = "Good Evening!"
	} else {
		message = "Good Night!"
	}

	return message, nil
}
