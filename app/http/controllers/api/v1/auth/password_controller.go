package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseAPIController
}

func (ctl *PasswordController) ResetByPhone(c *gin.Context) {
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}

	_user := user.GetByPhone(request.Phone)
	if _user.ID == 0 {
		response.Abort404(c)
	} else {
		_user.Password = request.Password
		_user.Save()

		response.Success(c)
	}
}
