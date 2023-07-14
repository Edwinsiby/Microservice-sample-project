package service

import (
	"context"
	"errors"
	"log"
	pb "service1/pb"
	"service1/pkg/entity"
	repo "service1/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type MyService struct {
	pb.UnimplementedMyServiceServer
}

func (s *MyService) MyMethod(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("Microservice1: MyMethod called")

	result := "Hello, " + req.Data
	return &pb.Response{Result: result}, nil
}

func (s *MyService) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.Response, error) {
	log.Println("Microservice1: Signup called")
	email, err := repo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("error with server")
	}
	if email != nil {
		return nil, errors.New("user with this email already exists")
	}
	phone, err := repo.GetByPhone(req.Phone)
	if err != nil {
		return nil, errors.New("error with server")
	}
	if phone != nil {
		return nil, errors.New("user with this phone no already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &entity.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  string(hashedPassword),
	}

	err1 := repo.Create(newUser)
	if err1 != nil {
		return nil, err1
	}

	result := "user created succesfuly"
	return &pb.Response{Result: result}, nil
}

func (s *MyService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := repo.GetByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user with this phone not found")
	}
	permission, err := repo.CheckPermission(user)
	if permission == false {
		return nil, errors.New("user permission denied")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("Invalid Password")
	} else {
		id := int32(user.ID)
		result := "user loged in succesfuly and cookie stored"
		return &pb.LoginResponse{Userid: id, Result: result}, nil
	}
}

func (s *MyService) AddAddress(ctx context.Context, req *pb.AddAddressRequest) (*pb.Response, error) {

	newAddress := &entity.Address{
		UserId:  int(req.Userid),
		House:   req.House,
		Street:  req.Street,
		City:    req.City,
		Pincode: req.Pincode,
		Type:    req.Type,
	}
	err := repo.CreateAddress(newAddress)
	if err != nil {
		return nil, err
	} else {
		result := "user address added succesfuly"
		return &pb.Response{Result: result}, nil
	}

}

func (s *MyService) AdminSignup(ctx context.Context, req *pb.AdminSignupRequest) (*pb.Response, error) {
	email, err := repo.AdminGetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if email != nil {
		return nil, errors.New("admin with this email already exists")
	}
	phone, err := repo.AdminGetByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	if phone != nil {
		return nil, errors.New("admin with this phone already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newAdmin := &entity.Admin{
		AdminName: req.Adminname,
		Email:     req.Email,
		Phone:     req.Phone,
		Role:      req.Role,
		Password:  string(hashedPassword),
	}

	err = repo.AdminCreate(newAdmin)
	if err != nil {
		return nil, err
	}

	result := "admin created succesfuly"
	return &pb.Response{Result: result}, nil
}

func (s *MyService) AdminLogin(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := repo.GetByPhone(req.Phone)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user with this phone not found")
	}
	permission, err := repo.CheckPermission(user)
	if permission == false {
		return nil, errors.New("user permission denied")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("Invalid Password")
	} else {
		id := int32(user.ID)
		result := "admin loged in succesfuly and cookie stored"
		return &pb.LoginResponse{Userid: id, Result: result}, nil
	}
}
