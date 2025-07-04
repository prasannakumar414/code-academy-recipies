package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SomeMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		safeHeader := c.Request.Header.Get("safe")
		if safeHeader != "" {
			fmt.Println("request is safe")
		}
		c.Next()
	}
}
