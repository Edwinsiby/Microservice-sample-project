package db

import (
	"errors"
	"log"
	"service2/pb"
	"service2/pkg/db"
	"service2/pkg/entity"

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

func ProductList() ([]entity.Apparel, error) {
	var apparels []entity.Apparel
	err := DB.Where("removed = ?", false).Find(&apparels).Error
	if err != nil {
		return nil, err
	}

	return apparels, nil

}

func GetApparelByID(id int) (*pb.Product, error) {
	var apparel entity.Apparel
	result := DB.First(&apparel, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
		return nil, result.Error
	}
	response := &pb.Product{
		Id:       int32(apparel.ID),
		Name:     apparel.Name,
		Price:    int32(apparel.Price),
		ImageURL: apparel.ImageURL,
		Category: apparel.Category,
	}
	return response, nil
}

func GetByUserID(userid int) (*entity.Cart, error) {
	var cart entity.Cart
	result := DB.Where("user_id=?", userid).First(&cart)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("cart not found")
		}
		return nil, errors.New("cart not found")
	}
	return &cart, nil
}

func Create(userid int) (*entity.Cart, error) {
	cart := &entity.Cart{
		UserId: userid,
	}
	if err := DB.Create(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil

}

func GetByName(productName string, cartId int) (*entity.CartItem, error) {
	var cartItem entity.CartItem
	result := DB.Where("product_name = ? AND cart_id = ?", productName, cartId).First(&cartItem)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &cartItem, nil
}

func CreateCartItem(cartItem *entity.CartItem) error {
	if err := DB.Create(cartItem).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCartItem(cartItem *entity.CartItem) error {
	return DB.Save(cartItem).Error
}
func UpdateCart(cart *entity.Cart) error {
	return DB.Where("user_id = ?", cart.UserId).Save(&cart).Error
}

func RemoveCartItem(cartItem *entity.CartItem) error {
	return DB.Where("product_name=?", cartItem.ProductName).Delete(&cartItem).Error
}
