package main

import (
	"flag"
	"fmt"
	"golang-application/application/repositories"
	"golang-application/application/usecases"
	"golang-application/framework/pb"
	"golang-application/framework/servers"
	"golang-application/framework/utils"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var db *gorm.DB

func main() {
	db = utils.ConnectDB()
	db.LogMode(true)

	port := flag.Int("port", 0, "Choose the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	userServer := setUpUserServer()

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}

func setUpUserServer() *servers.UserServer {
	userRepository := repositories.UserRepositoryDb{Db: db}
	userServer := servers.NewUserServer()
	userServer.UserUseCase = usecases.UserUseCase{UserRepository: &userRepository}
	return userServer
}
