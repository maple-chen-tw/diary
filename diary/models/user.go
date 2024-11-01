package models

import (
	"time"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *User) error {
	user.CreateAt = time.Now()
	return db.Create(user).Error
}

func GetUserByUsername(db *gorm.DB, username string) (User, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)
	return user, result.Error
}

func UpdateUser(db *gorm.DB, user *User) error {
	return db.Save(user).Error
}

func DeleteUser(db *gorm.DB, id int) error {
	return db.Delete(&User{}, id).Error
}
