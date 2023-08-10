package main

import (
	"flight/gin"
	"flight/handle"
	"log"
	"time"
)

func LogHook() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(LogHook())

	v1 := r.Group("/v1")
	{

		v1.POST("/tracker", handle.Track)
	}

	_ = r.Run("127.0.0.1:8080")
}
