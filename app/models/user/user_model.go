package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"gohub/pkg/hash"
)

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

func (user *User) Create() {
	database.DB.Create(&user)
}

func (user *User) ComparePassword(password string) bool {
	return hash.BcryptCheck(password, user.Password)
}

func (user *User) Save() int64 {
	result := database.DB.Save(&user)
	return result.RowsAffected
}
