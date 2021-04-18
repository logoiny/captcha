package router

import (
	"captcha/work"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)


func IpLimit() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := time.Now()
		reqIP := c.ClientIP()
		if reqIP == "::1" {
			reqIP = "127.0.0.1"
		}

		taste := time.Since(t)
		fmt.Println( "reqIP：", reqIP, "  ",taste )
		if work.IsIpLimit(reqIP){
			c.String(200, "The server can be accessed 30 times per minute")
		}
		c.Next()
	}
}