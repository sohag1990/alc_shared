package alc_shared

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// var API_URL = "http://localhost:8080/api"

// Cors set cors to routes
func Cors() gin.HandlerFunc {
	return cors.Default()

}
