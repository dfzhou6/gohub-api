package verifycode

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/helpers"
	"gohub/pkg/logger"
	"gohub/pkg/redis"
	"gohub/pkg/sms"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once

var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   config.GetString("app.name") + ":verifycode:",
			},
		}
	})

	return internalVerifyCode
}

func (vc *VerifyCode) generateVerifyCode(key string) string {
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}

	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})

	vc.Store.Set(key, code)
	return code
}

func (vc *VerifyCode) SendSMS(phone string) bool {
	code := vc.generateVerifyCode(phone)
	if !app.IsProduction() &&
		strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}

	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": code},
	})
}

func (vc *VerifyCode) CheckAns(key string, ans string) bool {
	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: ans})
	if !app.IsProduction() && (strings.HasSuffix(key, config.GetString("verifycode.debug_email_suffix")) ||
		strings.HasPrefix(key, config.GetString("verifycode.debug_phone_prefix"))) {
		return true
	}

	return vc.Store.Verify(key, ans, false)
}