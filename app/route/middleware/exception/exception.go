package exception

import (
	"../../../config"
	"../../../utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"strings"
	"time"
)

func SetUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				subject := fmt.Sprintf("[Panic - %s]项目出错了！", config.AppName)
				body := fmt.Sprintf("<b>错误时间: %s\n Runtime:\n<b>%s",
					time.Now().Format("2006/01/02 - 15:04:05"),
					string(debug.Stack()))

				bodyHtml := ""
				for _, v := range strings.Split(body, "\n") {
					bodyHtml += v + "<br>"
				}

				utilsGin := utils.Gin{Ctx: c}
				utilsGin.Response(500, subject, bodyHtml)
			}
		}()
		c.Next()
	}
}
