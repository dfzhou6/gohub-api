package user

import "gohub/pkg/database"

func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

func GetByPhone(phone string) (user User) {
	database.DB.Where("phone = ?", phone).First(&user)
	return
}

func GetByMulti(loginID string) (user User) {
	database.DB.Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("name = ?", loginID).
		First(&user)
	return
}

func Get(idStr string) (user User) {
	database.DB.Where("id", idStr).First(&user)
	return
}
