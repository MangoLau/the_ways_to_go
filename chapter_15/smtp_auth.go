package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"mangolau34@gmail.com",
		"@WOY1ZH1J1AND1NG",
		"smtp.gmail.com",
	)
	// Connect to the server, authentivate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:465",
		auth,
		"mangolau34@gmail.com",
		[]string{"mangolau34@gmail.com"},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}