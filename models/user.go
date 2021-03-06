package models

import (
	"github.com/mustanish/omelette/entity"
)

type User entity.User

//entity.DB.LogMode(true)

// Authorize model method
func (u *User) Authorize() (*User, error) {
	response := entity.DB.Create(&u)
	return u, response.Error
}

func (u *User) ReadDetail() {

}

func (u *User) UpdateDetail() error {
	response := entity.DB.Save(&u)
	return response.Error
}

func (u *User) Verify() {

}

// Check model method
func (u *User) Check(identity interface{}) *User {
	entity.DB.Raw("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND ((email = ?) OR (phone_number = ?) OR (user_name = ?) OR ID = ?)", identity, identity, identity, identity).Scan(&u)
	return u
}
