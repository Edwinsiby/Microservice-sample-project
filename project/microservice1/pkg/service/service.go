package service

import (
	"context"
	"log"
	pb "service1/pb"
)

type MyService struct {
	pb.UnimplementedMyServiceServer
}

// func init() {
// 	db, err := db.ConnectToDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func (s *MyService) MyMethod(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("Microservice1: MyMethod called")

	result := "Hello, " + req.Data
	return &pb.Response{Result: result}, nil
}

func (s *MyService) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.Response, error) {
	log.Println("Microservice1: Signup called")
	// email, err := db.GetByEmail(user.Email)
	// if err != nil {
	// 	return nil, errors.New("error with server")
	// }
	// if email != nil {
	// 	return nil, errors.New("user with this email already exists")
	// }
	// phone, err := db.userRepo.GetByPhone(user.Phone)
	// if err != nil {
	// 	return nil, errors.New("error with server")
	// }
	// if phone != nil {
	// 	return nil, errors.New("user with this phone no already exists")
	// }

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }

	// newUser := &entity.User{
	// 	FirstName: user.FirstName,
	// 	LastName:  user.LastName,
	// 	Email:     user.Email,
	// 	Phone:     user.Phone,
	// 	Password:  string(hashedPassword),
	// }

	// err1 := us.userRepo.Create(newUser)
	// if err1 != nil {
	// 	return nil, err1
	// }

	result := "Hello DBBB, "
	return &pb.Response{Result: result}, nil
}
