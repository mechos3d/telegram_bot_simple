package main

// var routingTable = map[string]string{}

// import (
// "fmt"
// )

func processMessage(msg string) (result string) {
	switch msg {
	case "/start":
		result = "you typed /start"
	default:
		result = msg
	}
	return result
}
