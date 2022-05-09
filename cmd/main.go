package main

import (
	"net"

	"github.com/amchicas/go-profile-srv/config"
	"github.com/amchicas/go-profile-srv/internal/domain"
	"github.com/amchicas/go-profile-srv/internal/service"
	"github.com/amchicas/go-profile-srv/internal/storage/mongo"
	"github.com/amchicas/go-profile-srv/pkg/log"
	"github.com/amchicas/go-profile-srv/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := log.New("Profile", "dev")
	c, err := config.LoadConfig()
	if err != nil {

		logger.Error("Failed at config" + err.Error())
	}
	repo := newMongo(c.MongoHost, c.MongoPort, c.Database)
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {

		logger.Error("Failed at server profile" + err.Error())
	}
	srv := service.New(repo, logger)
	grpcServer := grpc.NewServer()
	pb.RegisterProfileServiceServer(grpcServer, srv)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("Failed to server :" + err.Error())

	}
}

func newMongo(host, port, database string) domain.Repository {
	db, cancel := mongo.NewConn(host, port, database)
	defer cancel()
	return mongo.NewMongo(db)
}
