package db

import (
	"errors"
	"log"
	"service1/pkg/db"
	"service1/pkg/entity"

	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = db.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
}

func GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := DB.Where(&entity.User{Email: email}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
func GetByPhone(phone string) (*entity.User, error) {
	var user entity.User
	result := DB.Where(&entity.User{Phone: phone}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func Create(user *entity.User) error {
	return DB.Create(user).Error
}

func CheckPermission(user *entity.User) (bool, error) {
	result := DB.Where(&entity.User{Phone: user.Phone}).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	permission := user.Permission
	return permission, nil
}

func CreateAddress(address *entity.Address) error {
	return DB.Create(address).Error
}

func AdminCreate(admin *entity.Admin) error {
	return DB.Create(admin).Error
}

func AdminGetByPhone(phone string) (*entity.Admin, error) {
	var admin entity.Admin
	result := DB.Where(&entity.Admin{Phone: phone}).First(&admin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &admin, nil
}
func AdminGetByEmail(email string) (*entity.Admin, error) {
	var admin entity.Admin
	result := DB.Where(&entity.Admin{Email: email}).First(&admin)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &admin, nil
}
