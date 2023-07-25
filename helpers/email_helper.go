package helpers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/sohag1990/alc_shared/models"
)

// Company info
type CompanyInfo struct {
	CompanyName string
	Address     string
	Phone       string
	WhatsApp    string
	Website     string
	Email       string
	Logo        string
}

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

func RegistrationEmailTemplate(company CompanyInfo, user models.User) string {

	template := `
		<!DOCTYPE html>
		<html>

		<body style="font-family: Arial, sans-serif; line-height: 1.6;">

			<div style="max-width: 600px; margin: 0 auto; padding: 20px; border: 1px solid #e0e0e0; background-color: #f9f9f9;">
				<div style="text-align: center;">
					<img src="` + company.Logo + `" alt="` + company.CompanyName + ` Logo"
						style="width: 150px; height: auto;">
					<h1>` + company.CompanyName + `</h1>
				</div>

				<div style="margin-top: 20px;">
					<p>Hello ` + user.Profile.FullName + `,</p>
					<p>Thank you for registering on the <a target="_self" href="` + company.Website + `"> ` + company.CompanyName + ` website</a>. Your account has been successfully
						created.</p>
					<p>Please use the following credentials to log in:</p>
					<p><strong>Username/Email:</strong> ` + user.Email + `</p>
					<p><strong>Password:</strong> ` + user.Password + `</p>
					<p>We recommend keeping your login information secure and not sharing it with anyone.</p>
					<p>If you have any questions or need further assistance, please feel free to contact us.</p>
				</div>

				<div style="margin-top: 30px; text-align: center; font-size: 12px; color: #666666;">
					<p>` + company.CompanyName + `</p>
					<p` + company.Address + `</p>
					<p>Phone: ` + company.Phone + `</p>
					<p>WhatsApp: ` + company.WhatsApp + `</p>
					<p>Email: ` + company.Email + `</p>
					<p>Website: ` + company.Website + `</p>
				</div>
			</div>
		</body>

		</html>

		`
	return template
}
