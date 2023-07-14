package service

import (
	"context"
	"errors"
	"log"
	pb "service4/pb"
	"service4/pkg/entity"
	repo "service4/pkg/repository"
)

type MyService struct {
	pb.UnimplementedMyServiceServer
}

func (s *MyService) MyMethod(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("Microservice1: MyMethod called")

	result := "Hello, " + req.Data
	return &pb.Response{Result: result}, nil
}

func (s *MyService) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.Response, error) {
	err := repo.GetByApparelName(req.Name)
	if err == nil {
		return nil, errors.New("product already exists")
	}
	newApparel := &entity.Apparel{
		Name:     req.Name,
		Price:    int(req.Price),
		ImageURL: req.ImageURL,
		AdminId:  int(req.AdminId),
	}
	_, err1 := repo.CreateApparel(newApparel)
	if err1 != nil {
		return nil, err1
	} else {
		result := "Product Added Succesfuly"
		return &pb.Response{Result: result}, nil
	}
}

func (s *MyService) RemoveProduct(ctx context.Context, req *pb.RemoveProductRequest) (*pb.Response, error) {
	result, err := repo.GetApparelByID(int(req.Productid))
	if err != nil {
		return nil, err
	}
	result.Removed = !result.Removed
	err1 := repo.UpdateApparel(result)
	if err1 != nil {
		return nil, errors.New("Apparel deletion unsuccesfull")
	} else {
		result := "Product Removed Succesfuly"
		return &pb.Response{Result: result}, nil
	}
}
