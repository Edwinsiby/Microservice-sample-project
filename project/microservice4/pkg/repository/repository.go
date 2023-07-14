package repository

import (
	"errors"
	"log"
	"service4/pkg/db"
	"service4/pkg/entity"

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
func GetByApparelName(apparelName string) error {
	var existingApparel entity.Ticket
	result := DB.Where(&entity.Apparel{Name: apparelName}).First(&existingApparel)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
		return result.Error
	}
	return nil

}

func CreateApparel(apparel *entity.Apparel) (int, error) {
	if err := DB.Create(apparel).Error; err != nil {
		return 0, err
	}
	return apparel.ID, nil
}

func GetApparelByID(id int) (*entity.Apparel, error) {
	var apparel entity.Apparel
	result := DB.First(&apparel, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &apparel, nil
}
func UpdateApparel(apparel *entity.Apparel) error {
	return DB.Save(apparel).Error
}
