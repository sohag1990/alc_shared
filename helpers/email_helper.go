package helpers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sohag1990/alc_shared/models"
	"gopkg.in/gomail.v2"
)

func downloadFile(url, token, filename string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set the custom header "Token"
	if token != "" {
		req.Header.Set("Token", token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

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

func HMOFireRiskAssessmentCompany() CompanyInfo {
	var companyInfo CompanyInfo
	companyInfo.CompanyName = "HMO Fire Risk Assessment"
	companyInfo.Address = "55 Windermere Avenue, Purfleet-on-Thames RM19 1QN"
	companyInfo.Phone = "+44 20 8004 9655"
	companyInfo.WhatsApp = "+44 7856 377502"
	companyInfo.Email = "info@hmofireriskassessment.co.uk"
	companyInfo.Website = "https://hmofireriskassessment.co.uk"
	companyInfo.Logo = "https://hmofireriskassessment.co.uk/images/hmo-fire-risk-assessment-logo.png"

	return companyInfo
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
	AWSMTPUsername string
	AWSMTPPassword string
}

func (email Email) SendEmail(order models.Order, c *gin.Context) {

	token := jwt.GetToken(c)

	pdfURL := "http://localhost:8070/invoice-pdf/" + fmt.Sprint(order.ID)
	pdfFilename := fmt.Sprint(order.InvID) + ".pdf"
	if err := downloadFile(pdfURL, token, pdfFilename); err != nil {
		fmt.Println("Error downloading PDF:", err)
		return
	}

	// Compose the email
	m := gomail.NewMessage()
	m.SetHeader("From", email.FromEmail) // Replace with the sender email address
	m.SetHeader("To", email.ToEmails)    // Replace with the recipient email address
	m.SetHeader("Cc", email.CCEmails)    // Replace with your CC email address
	m.SetHeader("Bcc", email.BCCEmails)  // Replace with your BCC email address
	m.SetHeader("Subject", email.SubjectLine)
	m.SetBody("text/html", email.HtmlBody)

	// Attach the PDF file
	m.Attach(pdfFilename)

	// Send the email using SES
	d := gomail.NewDialer(email.AwsZone, 587, email.AWSMTPUsername, email.AWSMTPPassword) // Replace with your AWS credentials
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent through Amazon SES successfully!")
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
func GetServicesList(order models.Order) string {
	services := ""
	for _, service := range order.OrderServices {
		services += `
		<tr>
		<td>` + service.Title + `</td>
		<td>` + fmt.Sprint(service.Total) + `</td>
		</tr>
		`
	}
	return services
}
func NewOrderEmailTemplate(company CompanyInfo, user models.User, order models.Order) string {

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
					<h3>ORDER DETAILS: </h3>
					<p>Your order has been received and is now being processed. 
					Please feel free to email us on ` + company.Email + ` to get an instant support. 
					Your order details are shown below for your reference:</p>
					<h3>OUR BANK DETAILS</h3>
					<p>All Landlord Certificates Limited:</p>
					<ul>
					 <li>Account number: 0001 7807</li>
					 <li>Sort code: 30-93-44</li>
					</ul>
					<h3>ORDER #` + fmt.Sprint(order.InvID) + ` (` + order.CreatedAt.String() + `)</h3>
					<table>
						<thead>
							<tr>
								<th>Service</th>
								<th>Price</th>
							<tr>
						</thead>
						<tbody>` + GetServicesList(order) + `</tbody>
						<tfoot>
							<tr>
								<td>Subtotal:</td>
								<td>` + fmt.Sprint(order.Subtotal) + `</td>
							</tr>
							<tr>
								<td>Discount:</td>
								<td>` + fmt.Sprint(order.Discount) + `</td>
							</tr>
							<tr>
								<td>Total:</td>
								<td>` + fmt.Sprint(order.Total) + `</td>
							</tr>
						</tfoot>
					</table>
					<h3>BILLING ADDRESS</h3>
					<table>
						<tr><td>` + order.FullName + `</td></tr>
						<tr><td>` + order.CompanyName + `</td></tr>
						<tr><td>` + order.BillingAddress + ` ` + order.BillingPostCode + `</td></tr>
						<tr><td>` + order.Phone + `</td></tr>
						<tr><td>` + order.Email + `</td></tr>
					</table>

					<h3>Property ADDRESS</h3>
					<table>
						<tr><td>` + order.LandlordName + `</td></tr>
						<tr><td>` + order.PropertyAddress + ` ` + order.PropertyPostCode + `</td></tr>
						<tr><td>` + order.Phone + `</td></tr>
						<tr><td>` + order.Email + `</td></tr>
					</table>

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
