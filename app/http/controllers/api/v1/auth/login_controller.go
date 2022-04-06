package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/auth"
	"gohub/pkg/jwt"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIController
}

func (ctl *LoginController) LoginByPhone(c *gin.Context) {
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	_user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		response.Error(c, err, "账号不存在或密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}

}

func (ctl *LoginController) LoginByPassword(c *gin.Context) {
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	_user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		response.Unauthorized(c, "登录失败")
	} else {
		token := jwt.NewJWT().IssueToken(_user.GetStringID(), _user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

func (ctl *LoginController) RefreshToken(c *gin.Context) {
	token, err := jwt.NewJWT().RefreshToken(c)
	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
