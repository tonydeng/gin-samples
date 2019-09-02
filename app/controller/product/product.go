package product

import (
	"../../utils"
	"../param_bind"
	"../param_verify"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

func Add(c *gin.Context) {
	utilsGin := utils.Gin{Ctx: c}

	s, e := utils.Bind(&param_bind.ProductAdd{}, c)

	if e != nil {
		utilsGin.Response(-1, e.Error(), nil)
	}
	validate := validator.New()

	_ = validate.RegisterValidation("NameValid", param_verify.NameValid)

	if err:= validate.Struct(s);err !=nil {
		utilsGin.Response(-1,err.Error(),nil)

		return
	}

	utilsGin.Response(1,"success",nil)
}

// 编辑
func Edit(c *gin.Context) {
	fmt.Println(c.Request.RequestURI)
}

// 删除
func Delete(c *gin.Context) {
	fmt.Println(c.Request.RequestURI)
}

// 详情

func Detail(c *gin.Context) {
	fmt.Println(c.Request.RequestURI)
}