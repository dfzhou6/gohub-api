package hash

import (
	"gohub/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(password string) string {
	btyes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)
	return string(btyes)
}

func BcryptCheck(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func BcryptIsHashed(str string) bool {
	return len(str) == 60
}
