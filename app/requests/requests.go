package requests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors":  err.Error(),
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请用 JSON 格式",
		})
		fmt.Println(err.Error())
		return false
	}

	errs := handler(obj, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors":  errs,
			"message": "请求验证不通过，具体请查看 errors",
		})
		fmt.Println(errs)
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData,
	messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Messages:      messages,
		Rules:         rules,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}
