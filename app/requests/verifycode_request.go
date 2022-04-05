package requests

import (
	"gohub/pkg/captcha"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodePhoneRequest struct {
	CaptchaID  string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAns string `json:"captcha_ans,omitempty" valid:"captcha_ans"`
	Phone      string `json:"phone,omitempty" valid:"phone"`
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
	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaAns); !ok {
		errs["captcha_ans"] = append(errs["captcha_ans"], "图片验证码错误")
	}

	return errs
}
