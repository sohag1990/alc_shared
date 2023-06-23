package helpers

import (
	"log"
	"net/smtp"
)

func main() {
	// Sender's email credentials
	senderEmail := "your_email@example.com"
	senderPassword := "your_email_password"

	// Recipient's email
	recipientEmail := "recipient@example.com"

	// SMTP server configuration
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	// Compose the email
	subject := "Hello from Go!"
	body := "This is the email body."

	// Create the message
	message := "From: " + senderEmail + "\n" +
		"To: " + recipientEmail + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Authenticate and send the email
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{recipientEmail}, []byte(message))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")
}
