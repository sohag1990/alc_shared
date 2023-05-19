package helpers

import (
	"encoding/json"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sohag1990/alc_shared/models"
)

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Hostname string `form:"hostname" json:"hostname" binding:"required"`
}

var identityKey = "id"
var fullNameKey = "fullname"
var emailKey = "email"
var roleKey = "role"
var hostnameKey = "hostname"

type LogedInUser struct {
	ID       uint64
	Email    string
	FullName string
	Role     string
	Hostname string
}

// GinJwtMiddlewareHandler to handle authentication and authorization for api routes
func GinJwtMiddlewareHandler() *jwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("110481"),
		Timeout:     time.Hour * 720,
		MaxRefresh:  time.Hour * 720,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*LogedInUser); ok {
				return jwt.MapClaims{
					identityKey: v.ID,       // identity key is user ID
					fullNameKey: v.FullName, // Name
					emailKey:    v.Email,    // Email
					roleKey:     v.Role,     // Role
					hostnameKey: v.Hostname, // Hostname
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			return &LogedInUser{
				ID:       uint64(claims[identityKey].(float64)), // identity key is user ID
				Email:    claims[emailKey].(string),
				FullName: claims[fullNameKey].(string),
				Role:     claims[roleKey].(string),
				Hostname: claims[hostnameKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			apiUrl := "http://localhost:8080/api/check_login"
			var usr models.User
			usr.Email = loginVals.Email
			usr.Password = loginVals.Password
			result := APICall("POST", apiUrl, usr, usr, c)
			if err := json.Unmarshal(result, &usr); err != nil {
				panic(err)
			}

			if usr.ID == 0 {
				c.AbortWithStatus(401)
				return nil, jwt.ErrFailedAuthentication
			} else {
				return &LogedInUser{
					ID:       usr.ID,
					Email:    usr.Email,
					FullName: usr.Profile.FullName,
					Role:     usr.Role,
					Hostname: loginVals.Hostname,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*LogedInUser); ok && v.Role != "" { // here you can set a great logic
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}

func APIUserLoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		claims := jwt.ExtractClaims(c)
		UserRole := claims["role"]
		if UserRole == "admin" || UserRole == "customer" {
			c.Next()
		} else {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "Unauthenticated",
			})
		}
	}
}

func APIAdminLoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		claims := jwt.ExtractClaims(c)
		UserRole := claims["role"]
		if UserRole == "admin" {
			c.Next()
		} else {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "Unauthenticated",
			})
		}
	}
}
