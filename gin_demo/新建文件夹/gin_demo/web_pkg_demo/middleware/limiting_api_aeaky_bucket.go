package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

/**
漏桶限流器
*/
var limitChanLeakyBucket chan *gin.Context

func ApiInLeakyBucket(c *gin.Context) error {
	// 入桶
	if len(limitChanLeakyBucket) < cap(limitChanLeakyBucket) {
		limitChanLeakyBucket <- c
		return nil
	}
	return errors.New("桶已满，请求被放弃")
}

// ApiOutLeakyBucket 漏桶限流的方法对出桶没有要求，直接放行
func ApiOutLeakyBucket(c *gin.Context) {
	c.Next()
}

// LimitingApiLeakyBucket 入桶检测错误，出桶直接放行，桶容量为100
func LimitingApiLeakyBucket() gin.HandlerFunc {
	// 桶大小100
	limitChanLeakyBucket = make(chan *gin.Context, 100)

	// 起一个定时器，每隔固定时间从桶中获取请求进行处理
	go tickerLeakyBucket()
	return func(c *gin.Context) {
		// 入桶，若失败，则有错误信息
		err := ApiInLeakyBucket(c)
		if err != nil {
			// 入桶失败，友好处理
			// fixme: 错误返回后的具体处理，需要再构思
			c.AbortWithError(500, err)
		}

	}
}

// 计时器
func tickerLeakyBucket() {
	// 10ms 取出一个请求
	ticker := time.NewTicker(time.Millisecond * 10)
	// 在合适的地方重新设置漏水速率
	//ticker.Reset(time.Microsecond*100)
	for {
		<-ticker.C
		// 本次是否取出
		if len(limitChanLeakyBucket) > 0 {
			// fixme: 取出处理
			c := <-limitChanLeakyBucket
			ApiOutLeakyBucket(c)
		}
		//fmt.Println("tick:", i)
	}
}
