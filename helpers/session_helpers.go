package helpers

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	jwt_go "github.com/golang-jwt/jwt/v4"
	"github.com/sohag1990/alc_shared/models"
)

func GetFrontendUserData(c *gin.Context) models.FrontendUserData {

	// token key hishabe use kore session create kora jai kina dekho. tahole secure hobe
	session := sessions.Default(c)

	tokenString := fmt.Sprint(session.Get("login_session"))
	claims := ClaimsToken(tokenString)

	var usrData models.FrontendUserData
	if claims != nil || len(claims) != 0 {
		usrData.FullName = fmt.Sprint(claims["fullname"])
		usrData.Email = fmt.Sprint(claims["email"])
		usrData.Role = fmt.Sprint(claims["role"])
		usrData.Hostname = fmt.Sprint(claims["hostname"])
		usrData.ButtonText = "Logout"
		usrData.Url = "/logout"
		return usrData
	}
	usrData.FullName = "Guest"
	usrData.Email = ""
	usrData.Role = "customer"
	usrData.ButtonText = "Login"
	usrData.Url = "/login"

	return usrData

}

func ClaimsToken(tokenString string) jwt_go.MapClaims {
	//jwt_go "github.com/golang-jwt/jwt/v4" library dependancy
	token, err := jwt_go.Parse(tokenString, func(token *jwt_go.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt_go.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("110481"), nil // app secrate
	})
	if err != nil {
		return nil
	}
	claims, ok := token.Claims.(jwt_go.MapClaims)
	if ok && token.Valid {

		return claims
	}
	return nil
}
