// ORM-модель User
package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint32 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
