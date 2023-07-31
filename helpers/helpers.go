package helpers

import (
	"fmt"
	"html/template"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sohag1990/alc_shared/models"
)

func InvIDGenerator(order models.Order) string {

	invID := GetYearFromTime(order.CreatedAt) + "-" + Uint64ConvertToString(order.InvID)
	return invID
}
func StringToInt(str string) int {
	intVal, _ := strconv.Atoi(str)
	return intVal
}
func Uint64ConvertToString(value uint64) string {
	return strconv.FormatUint(value, 10)
}

func GetYearFromTime(t *time.Time) string {
	if t != nil {
		return t.Format("2006")
	}
	return ""
}
func GetGetFormattedDate(t *time.Time) string {
	if t != nil {
		return t.Format("1/2/2006")
	}
	return ""

}
func GetGetFormattedInvDate(order models.Order) string {
	if order.CreatedAt != nil {
		return order.CreatedAt.Format("1/2/2006")
	}
	return ""
}
func Noescape(str string) template.HTML {
	return template.HTML(str)
}

func TokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// GeneratePassword generates a random password of the specified length and complexity level.
// Complexity levels are 1, 2, and 3, with 1 being the simplest and 3 being the most complex.
func GeneratePassword(length, complexity int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numberRunes = []rune("0123456789")
	var symbolRunes = []rune("!@#$%^&*()_+{}[]|\\;:'\"<>,.?/")

	rand.Seed(time.Now().UnixNano())

	var password []rune
	for i := 0; i < length; i++ {
		switch complexity {
		case 1:
			password = append(password, letterRunes[rand.Intn(len(letterRunes))])
		case 2:
			if rand.Intn(2) == 0 {
				password = append(password, letterRunes[rand.Intn(len(letterRunes))])
			} else {
				password = append(password, numberRunes[rand.Intn(len(numberRunes))])
			}
		case 3:
			switch rand.Intn(3) {
			case 0:
				password = append(password, letterRunes[rand.Intn(len(letterRunes))])
			case 1:
				password = append(password, numberRunes[rand.Intn(len(numberRunes))])
			case 2:
				password = append(password, symbolRunes[rand.Intn(len(symbolRunes))])
			}
		}
	}

	return string(password)
}

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// IsValidEmailFormat Email invalid formate checker
func IsValidEmailFormat(email string) bool {
	if !emailRegexp.MatchString(email) {
		return false
	}
	return true
}

func UrlSchema(c *gin.Context) string {
	isHTTPS := c.Request.TLS != nil

	// Determine the URL schema based on the HTTPS flag
	scheme := "http"
	if isHTTPS {
		scheme = "https"
	} else {
		scheme = "http"
	}
	return scheme
}
func GetHosturl(c *gin.Context) string {
	return UrlSchema(c) + "://" + c.Request.Host
}
func GetHostWithoutProtocol(c *gin.Context) string {
	return c.Request.Host
}
