package gateway

import (
	"context"
	"log"
	"net/http"

	"gateway/pb"

	"google.golang.org/grpc"
)

func NewAPIGateway() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Check the path and route the request accordingly
		switch r.URL.Path {
		case "/api/service1":
			conn1, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Failed to connect to service1: %v", err)
			}
			defer conn1.Close()
			client := pb.NewMyServiceClient(conn1)
			req := &pb.Request{
				Data: "Mydata",
			}
			resp, err := client.MyMethod(context.Background(), req)
			if err != nil {
				log.Fatalf("Failed to call MyMethod: %v", err)
			}
			w.Write([]byte(resp.Result))
		case "/api/service1/signup":
			conn1, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Failed to connect to service1: %v", err)
			}
			defer conn1.Close()
			client := pb.NewMyServiceClient(conn1)
			req := &pb.SignupRequest{
				Firstname: "Edwin",
				Lastname:  "Siby",
				Email:     "edwin@gmail.com",
				Phone:     "9048402133",
				Password:  "pass@123",
			}
			resp, err := client.Signup(context.Background(), req)
			if err != nil {
				log.Fatalf("Failed to call MyMethod: %v", err)
			}
			w.Write([]byte(resp.Result))
		case "/api/service2":
			conn2, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("Failed to connect to service1: %v", err)
			}
			defer conn2.Close()

		case "/api/get":

			w.Write([]byte("This is the GET endpoint"))
		default:
			http.NotFound(w, r)
		}
	})
}
