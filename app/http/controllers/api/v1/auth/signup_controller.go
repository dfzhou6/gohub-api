package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type SignUpController struct {
	v1.BaseAPIController
}

func (ctl *SignUpController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}

	if !requests.Validate(c, &request, requests.ValidateSignupPhoneExist) {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func (ctl *SignUpController) IsEmailExist(c *gin.Context) {

	request := requests.SignupEmailExistRequest{}

	if !requests.Validate(c, &request, requests.ValidateSignupEmailExist) {
		return
	}

	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
