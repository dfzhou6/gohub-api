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
			authGroup.POST("/signup/using-email", signUpCtl.SignupUsingEmail)

			verifyCodeCtl := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", verifyCodeCtl.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", verifyCodeCtl.SendUsingPhone)
			authGroup.POST("/verify-codes/email", verifyCodeCtl.SendUsingEmail)

			loginCtl := new(auth.LoginController)
			authGroup.POST("/login/using-phone", loginCtl.LoginByPhone)
			authGroup.POST("/login/using-password", loginCtl.LoginByPassword)
			authGroup.POST("/login/refresh-token", loginCtl.RefreshToken)

			pwdCtl := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", pwdCtl.ResetByPhone)
			authGroup.POST("/password-reset/using-email", pwdCtl.ResetByEmail)
		}

	}
}
