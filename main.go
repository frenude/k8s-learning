package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	r.GET("/", func(c *gin.Context) {
		ip := c.ClientIP()
		hostName, _ := os.Hostname()
		c.JSON(200, gin.H{
			"ip":       ip,
			"hostname": hostName,
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hostname", func(c *gin.Context) {
		hostName, _ := os.Hostname()
		c.JSON(200, gin.H{
			"message": hostName,
		})
	})
	r.MaxMultipartMemory = 64 << 20 // 64 MiB

	r.Run(":80")
}
