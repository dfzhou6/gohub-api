package routes

import (
	"gohub/app/http/controllers/api/v1/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World",
			})
		})

		authGroup := v1.Group("auth")
		{
			signUpCtl := new(auth.SignUpController)
			authGroup.POST("/signup/phone/exist", signUpCtl.IsPhoneExist)
			authGroup.POST("/signup/email/exist", signUpCtl.IsEmailExist)
			authGroup.POST("/signup/using-phone", signUpCtl.SignupUsingPhone)

			verifyCodeCtl := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", verifyCodeCtl.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", verifyCodeCtl.SendUsingPhone)
			authGroup.POST("/verify-codes/email", verifyCodeCtl.SendUsingEmail)
		}

	}
}
