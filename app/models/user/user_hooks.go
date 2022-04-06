package user

import (
	"gohub/pkg/hash"

	"gorm.io/gorm"
)

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(user.Password) {
		user.Password = hash.BcryptHash(user.Password)
	}
	return
}
