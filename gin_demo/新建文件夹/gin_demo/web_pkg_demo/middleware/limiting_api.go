package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

var limitChan chan int

func ApiIn(c *gin.Context) error {
	if len(limitChan) == cap(limitChan) {
		//fmt.Println("太忙了，限流了")
		return errors.New("太忙了，限流了")
	}
	limitChan <- 1
	return nil
}
func ApiOut(c *gin.Context) {
	<-limitChan
}

func LimitingApi() gin.HandlerFunc {
	limitChan = make(chan int, 100)
	fmt.Printf("限流容量：%d,当前限流大小：%d\n", cap(limitChan), len(limitChan))
	return func(c *gin.Context) {

		if err := ApiIn(c); err != nil {
			c.AbortWithError(500, err)
			//c.Abort()
			return
		}
		defer ApiOut(c)
		c.Next()
	}
}
