package helpers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Email struct {
	HtmlBody       string
	SubjectLine    string
	FromEmail      string
	ToEmails       string
	CCEmails       string
	BCCEmails      string
	AwsZone        string
	AWSAccessKeyID string
	AWSSecretKeyID string
}

func (email Email) SendEmail() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(email.AwsZone), // Replace with your desired AWS region
		Credentials: credentials.NewStaticCredentials(
			email.AWSAccessKeyID, // Replace with your AWS access key ID
			email.AWSSecretKeyID, // Replace with your AWS secret access key
			"",
		),
	})
	if err != nil {
		fmt.Println("Error creating AWS session:", err)
	}
	// Create an SES client
	svc := ses.New(sess)

	// Compose the email with an HTML body and CC address
	toAddresses := []*string{aws.String(email.ToEmails)}  // Replace with your to email address
	ccAddresses := []*string{aws.String(email.CCEmails)}  // Replace with your CC email address
	bcAddresses := []*string{aws.String(email.BCCEmails)} // Replace with your bCC email address

	input := &ses.SendEmailInput{
		Source: aws.String(email.FromEmail), // Replace with the sender email address
		Destination: &ses.Destination{
			ToAddresses:  toAddresses,
			CcAddresses:  ccAddresses,
			BccAddresses: bcAddresses,
		},
		Message: &ses.Message{
			Subject: &ses.Content{
				Data: aws.String(email.SubjectLine),
			},
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(email.HtmlBody),
				},
			},
		},
	}

	// Send the email through Amazon SES
	result, err := svc.SendEmail(input)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent! Message ID:", *result.MessageId)

}
