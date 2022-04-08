package routes

import (
	v1Ctrl "gohub/app/http/controllers/api/v1"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("v1")
	v1.Use(middlewares.LimitIP("100-H"))
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World",
			})
		})

		userCtrl := new(v1Ctrl.UsersController)
		v1.GET("/user", middlewares.AuthJWT(), userCtrl.CurrentUser)
		userGroup := v1.Group("/users")
		{
			userGroup.GET("", userCtrl.Index)
		}

		cateCtrl := new(v1Ctrl.CategoriesController)
		cateGroup := v1.Group("/categories")
		{
			cateGroup.POST("", middlewares.AuthJWT(), cateCtrl.Store)
			cateGroup.PUT("/:id", middlewares.AuthJWT(), cateCtrl.Update)
		}

		authGroup := v1.Group("auth")
		authGroup.Use(middlewares.LimitIP("200-H"))
		{
			signUpCtl := new(auth.SignUpController)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), signUpCtl.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), signUpCtl.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), signUpCtl.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), signUpCtl.SignupUsingEmail)

			verifyCodeCtl := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), verifyCodeCtl.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), verifyCodeCtl.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), verifyCodeCtl.SendUsingEmail)

			loginCtl := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), loginCtl.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), loginCtl.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), loginCtl.RefreshToken)

			pwdCtl := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwdCtl.ResetByPhone)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwdCtl.ResetByEmail)
		}

	}
}
