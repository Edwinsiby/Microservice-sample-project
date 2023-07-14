package service

import (
	"context"
	"errors"
	"log"

	pb "service2/pb"
	"service2/pkg/entity"
	repo "service2/pkg/repository"
)

type MyService struct {
	pb.UnimplementedMyServiceServer
}

func (s *MyService) MyMethod(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("Microservice2: MyMethod called")

	result := "Hola, " + req.Data
	return &pb.Response{Result: result}, nil
}

func (s *MyService) ProductList(ctx context.Context, req *pb.Request) (*pb.ProductListResponse, error) {
	apparelList, err := repo.ProductList()
	if err != nil {
		return nil, err
	}
	productList := ConvertToProtoProductList(apparelList)
	return productList, nil
}

func ConvertToProtoProductList(apparelList []entity.Apparel) *pb.ProductListResponse {
	productList := &pb.ProductListResponse{}
	for _, apparel := range apparelList {
		product := &pb.Product{
			Id:       int32(apparel.ID),
			Name:     apparel.Name,
			Price:    int32(apparel.Price),
			ImageURL: apparel.ImageURL,
			Category: apparel.Category,
		}

		productList.Apparels = append(productList.Apparels, product)
	}
	return productList
}

func (s *MyService) ProductDetails(ctx context.Context, req *pb.ProductDetailsRequest) (*pb.ProductDetailsResponse, error) {
	apparel, err := repo.GetApparelByID(int(req.Productid))
	if err != nil {
		return nil, err
	}
	response := &pb.ProductDetailsResponse{
		Apparel: apparel,
	}
	return response, nil
}

func (s *MyService) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.Response, error) {
	var userCart *entity.Cart
	var cartId int
	userCart, err := repo.GetByUserID(int(req.Userid))
	if err != nil {
		cart, err1 := repo.Create(int(req.Userid))
		if err1 != nil {
			return nil, errors.New("Failed to create user cart")
		}
		userCart = cart
		cartId = int(cart.ID)
	} else {
		cartId = int(userCart.ID)
	}
	apparel, err := repo.GetApparelByID(int(req.Productid))
	if err != nil {
		return nil, errors.New("Apparel not found")
	}
	cartItem := &entity.CartItem{
		CartId:      cartId,
		ProductId:   int(apparel.Id),
		Category:    "apparel",
		Quantity:    1,
		ProductName: apparel.Name,
		Price:       float64(apparel.Price),
	}
	existingApparel, err := repo.GetByName(apparel.Name, cartId)
	if existingApparel == nil {
		err = repo.CreateCartItem(cartItem)
		if err != nil {
			return nil, errors.New("Adding new apparel to cart item failed")
		}
	} else {
		existingApparel.Quantity += 1
		err := repo.UpdateCartItem(existingApparel)
		if err != nil {
			return nil, errors.New("error updating existing cartitem")
		}
	}
	userCart.TotalPrice += cartItem.Price * float64(1)
	userCart.ApparelQuantity += 1
	err = repo.UpdateCart(userCart)
	if err != nil {
		return nil, errors.New("Cart price updation failed")
	}

	result := "Product Added"
	return &pb.Response{Result: result}, nil
}

func (s *MyService) RemoveFromCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.Response, error) {
	userCart, err := repo.GetByUserID(int(req.Userid))
	if err != nil {
		return nil, errors.New("Failed to find user cart")
	}
	apparel, err := repo.GetApparelByID(int(req.Productid))
	if err != nil {
		return nil, errors.New("Apparel not found")
	}
	existingApparel, err1 := repo.GetByName(apparel.Name, int(userCart.ID))
	if err1 != nil {
		return nil, errors.New("Removing apparel from cart failed")
	}
	if existingApparel.Quantity == 1 {
		err := repo.RemoveCartItem(existingApparel)
		if err != nil {
			return nil, errors.New("Removin apparel from cart failed")
		}
	} else {
		existingApparel.Quantity -= 1
		err := repo.UpdateCartItem(existingApparel)
		if err != nil {
			return nil, errors.New("error updating existing cartitem")
		}
	}
	userCart.TotalPrice -= float64(apparel.Price)
	userCart.ApparelQuantity -= 1
	if userCart.OfferPrice > 0 {
		userCart.OfferPrice = 0
	}
	err = repo.UpdateCart(userCart)
	if err != nil {
		return nil, errors.New("Remove from cart failed")
	}

	result := "Product Removed"
	return &pb.Response{Result: result}, nil
}

func (s *MyService) CartDetails(ctx context.Context, req *pb.CartDetailsRequest) (*pb.CartDetailsResponse, error) {
	userCart, err := repo.GetByUserID(int(req.Userid))
	if err != nil {
		return nil, errors.New("Failed to find user cart")
	} else {

		response := &pb.CartDetailsResponse{
			UserId:          int32(userCart.UserId),
			Apparelquantity: int32(userCart.ApparelQuantity),
			Totalprice:      int32(userCart.TotalPrice),
			Offerprice:      int32(userCart.OfferPrice),
		}
		return response, nil
	}
}
