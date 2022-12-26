package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
)

func main() {
	wd, er := os.Getwd()
	if er != nil {
		fmt.Println("getwd error")
		fmt.Println(er)
	}
	fmt.Printf("Working dir = %v\n", wd)

	// Sender data.
	from := os.Getenv("FROM")
	password := os.Getenv("PASSWORD")"

	// Receiver email address.
	to := []string{
		"gsotoreyes37@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// t, _ := template.ParseFiles("template.html")

	t, _ := template.ParseFiles("template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Reference Letter Submission \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		RecipinentName string
		SenderName    string
		Company  string
	}{
		RecipinentName: "Goose",
		SenderName: "Bill Gates",
		Company:  "GitHub",
	})

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent!")
}
