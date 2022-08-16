package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

/**
令牌桶限流器
fixme: 三个限流策略封装成工具
*/

var limitChanTokenBucket chan bool

func ApiInTokenBucket(c *gin.Context) error {
	// 取出一张票
	if x := <-limitChanTokenBucket; x {

	}
	return nil
}
func ApiOutTokenBucket(c *gin.Context) {
}

func LimitingApiTokenBucket() gin.HandlerFunc {
	limitChanTokenBucket = make(chan bool, 10)
	for i := 0; i < cap(limitChanTokenBucket); i++ {
		limitChanTokenBucket <- true
	}
	// 协程取放票
	go tickerTokenBucket()
	fmt.Printf("限流容量：%d,当前限流大小：%d\n", cap(limitChanTokenBucket), len(limitChanTokenBucket))
	// 起一个定时器，不断放令牌，满了不放

	return func(c *gin.Context) {
		if len(limitChanTokenBucket) == 0 {
			c.AbortWithError(500, errors.New("繁忙"))
		} else {
			<-limitChanTokenBucket
			c.Next()
		}
	}
}

// 计时器
func tickerTokenBucket() {
	// 10ms放一张票
	ticker := time.NewTicker(time.Millisecond * 10)
	// 在合适的地方重新设置放票速率
	//ticker.Reset(time.Microsecond*100)
	i := 0
	for {
		<-ticker.C
		// 放一张票
		if len(limitChanTokenBucket) < cap(limitChanTokenBucket) {
			limitChanTokenBucket <- true
		}
		i++
		//fmt.Println("tick:", i)
	}
}
