package middleware

import (
	"fmt"
	"gin_demo/web_pkg_demo/lib"
	"github.com/gin-gonic/gin"
)

// PreHeader 预处理请求头,给后面的中间件使用
func PreHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.GetHeader("db")
		switch db {
		case "mysql":
		case "postgres":
			c.Set("db", lib.DBMapPool[db])
			c.Set("gdb", lib.GORMMapPool[db])
			break
		default:
			fmt.Println("数据库不支持!!这里需要再做处理")
		}
		c.Next()
	}
}
