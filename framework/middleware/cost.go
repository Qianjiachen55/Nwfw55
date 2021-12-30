package middleware

import (
	"github.com/Qianjiachen55/Nwfw55/framework/gin"
	"log"
	"time"
)

func Cost() gin.HandlerFunc {

	return func(c *gin.Context)  {
		start := time.Now()

		c.Next()

		end := time.Now()
		cost := end.Sub(start)

		log.Printf("api uri: %v, cost: %v", c.Request.RequestURI,cost.Seconds())

	}
}
