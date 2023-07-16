package servers

import (
	"context"
	"golang-application/application/usecases"
	"golang-application/domain"
	"golang-application/framework/pb"
	"log"
)

type UserServer struct {
	User        domain.User
	UserUseCase usecases.UserUseCase
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (userServer *UserServer) CreateUser(ctx context.Context, request *pb.UserRequest) (*pb.UserResponse, error) {
	userServer.User.Name = request.GetName()
	userServer.User.Email = request.GetEmail()
	userServer.User.Password = request.GetPassword()

	user, err := userServer.UserUseCase.Create(&userServer.User)

	if err != nil {
		log.Fatalf("Error during the RPC Create User: %v", err)
	}

	return &pb.UserResponse{Token: user.Token}, nil
}
