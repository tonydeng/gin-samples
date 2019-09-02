package route

import (
	"../controller/product"
	"../utils"
	"./middleware/exception"
	"./middleware/logger"
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {
	// set router

	engine.Use(logger.SetUp(), exception.SetUp())

	engine.NoRoute(func(c *gin.Context) {
		utilsGin := utils.Gin{Ctx: c}
		utilsGin.Response(404, "当前URI不存在", nil)
	})

	engine.GET("/ping", func(c *gin.Context) {
		utilsGin := utils.Gin{Ctx: c}
		utilsGin.Response(1, "pong", nil)
	})

	ProductRouter := engine.Group("/product")
	{
		ProductRouter.POST("",product.Add)

		// 更新产品
		ProductRouter.PUT("/:id", product.Edit)

		// 删除产品
		ProductRouter.DELETE("/:id", product.Delete)

		// 获取产品详情
		ProductRouter.GET("/:id", product.Detail)
	}
}
