package validators

import (
	"gohub/pkg/captcha"
	"gohub/pkg/verifycode"
)

func ValidateCaptcha(captchaID string, captchaAns string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAns); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}

func ValidatePasswordConfirm(password, passwordConfirm string, errs map[string][]string) map[string][]string {
	if password != passwordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配")
	}
	return errs
}

func ValidateVerifyCode(key, ans string, errs map[string][]string) map[string][]string {
	if ok := verifycode.NewVerifyCode().CheckAns(key, ans); !ok {
		errs["verify_code"] = append(errs["verify_code"], "验证码错误")
	}
	return errs
}
