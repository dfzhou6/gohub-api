package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
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
