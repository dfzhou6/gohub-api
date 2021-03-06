package routes

import (
	v1Ctrl "gohub/app/http/controllers/api/v1"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"
	"gohub/pkg/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("api/v1")
	} else {
		v1 = r.Group("v1")
	}

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
			userGroup.PUT("", middlewares.AuthJWT(), userCtrl.UpdateProfile)
			userGroup.PUT("/email", middlewares.AuthJWT(), userCtrl.UpdateEmail)
			userGroup.PUT("/phone", middlewares.AuthJWT(), userCtrl.UpdatePhone)
			userGroup.PUT("/password", middlewares.AuthJWT(), userCtrl.UpdatePassword)
			userGroup.PUT("/avatar", middlewares.AuthJWT(), userCtrl.UpdateAvatar)
		}

		cateCtrl := new(v1Ctrl.CategoriesController)
		cateGroup := v1.Group("/categories")
		{
			cateGroup.GET("", cateCtrl.Index)
			cateGroup.POST("", middlewares.AuthJWT(), cateCtrl.Store)
			cateGroup.PUT("/:id", middlewares.AuthJWT(), cateCtrl.Update)
			cateGroup.DELETE("/:id", middlewares.AuthJWT(), cateCtrl.Delete)
		}

		topicCtrl := new(v1Ctrl.TopicsController)
		topicGroup := v1.Group("/topics")
		{
			topicGroup.GET("", topicCtrl.Index)
			topicGroup.POST("", middlewares.AuthJWT(), topicCtrl.Store)
			topicGroup.PUT("/:id", middlewares.AuthJWT(), topicCtrl.Update)
			topicGroup.DELETE("/:id", middlewares.AuthJWT(), topicCtrl.Delete)
			topicGroup.GET("/:id", topicCtrl.Show)
		}

		linkCtrl := new(v1Ctrl.LinksController)
		linkGroup := v1.Group("/links")
		{
			linkGroup.GET("", linkCtrl.Index)
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
