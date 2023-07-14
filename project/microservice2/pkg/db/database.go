package db

import (
	"fmt"
	"service2/pkg/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var KEY5 = "host=localhost user=edwin dbname=microservice password=acid port=5432 sslmode=disable"

func ConnectToDB() (*gorm.DB, error) {
	// config, err := utils.LoadConfig("./")
	db, err := gorm.Open(postgres.Open(KEY5), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	DB = db
	DB.AutoMigrate(&entity.TicketDetails{}, &entity.OtpKey{}, &entity.Signup{}, &entity.Admin{}, &entity.User{}, &entity.Ticket{}, &entity.Apparel{}, &entity.CartItem{}, &entity.Cart{}, &entity.Wishlist{}, &entity.Order{}, &entity.OrderItem{}, &entity.Address{}, &entity.Inventory{}, &entity.Invoice{}, &entity.Return{}, &entity.Coupon{}, &entity.UsedCoupon{}, &entity.Offer{})
	return db, nil
}