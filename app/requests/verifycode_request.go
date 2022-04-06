package requests

import (
	"gohub/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodePhoneRequest struct {
	CaptchaID  string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAns string `json:"captcha_ans,omitempty" valid:"captcha_ans"`
	Phone      string `json:"phone,omitempty" valid:"phone"`
}

type VerifyCodeEmailRequest struct {
	CaptchaID  string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAns string `json:"captcha_ans,omitempty" valid:"captcha_ans"`
	Email      string `json:"email,omitempty" valid:"email"`
}

func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"captcha_id":  []string{"required"},
		"captcha_ans": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数为 phone",
			"digits:手机号的长度必须为 11 位数字",
		},
		"captcha_id": []string{
			"required:图片验证码ID为必填项",
		},
		"captcha_ans": []string{
			"required:图片验证码答案为必填项",
			"digits:图片验证码长度必须是 6 位数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*VerifyCodePhoneRequest)

	return validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAns, errs)
}

func VerifyCodeEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":       []string{"required", "min:4", "max:30", "email"},
		"captcha_id":  []string{"required"},
		"captcha_ans": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*VerifyCodeEmailRequest)
	return validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAns, errs)
}
