package user

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

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

func GetByEmail(email string) (user User) {
	database.DB.Where("email", email).First(&user)
	return
}

func All() (users []User) {
	database.DB.Find(&users)
	return
}

func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}
