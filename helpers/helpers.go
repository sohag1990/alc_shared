package helpers

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"strconv"
	"time"

	"github.com/sohag1990/alc_shared/models"
)

func InvIDGenerator(companyShortName string, order models.Order) string {
	fmt.Println(order.CreatedAt)
	invID := companyShortName + "-" + GetYearFromTime(order.CreatedAt) + "-" + Uint64ConvertToString(order.InvID)
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
