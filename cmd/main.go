package main

import (
	"context"

	"github.com/amchicas/go-profile-srv/config"
	"github.com/amchicas/go-profile-srv/internal/adder"
	"github.com/amchicas/go-profile-srv/internal/domain"
	"github.com/amchicas/go-profile-srv/internal/eraser"
	"github.com/amchicas/go-profile-srv/internal/fetcher"
	"github.com/amchicas/go-profile-srv/internal/grpc"
	"github.com/amchicas/go-profile-srv/internal/modifier"
	"github.com/amchicas/go-profile-srv/internal/storage/mongo"
	"github.com/amchicas/go-profile-srv/pkg/log"
	"golang.org/x/sync/errgroup"
)

func main() {
	logger := log.New("Profile", "dev")
	c, err := config.LoadConfig()
	if err != nil {

		logger.Error("Failed at config" + err.Error())
	}
	repo := newMongo(c.MongoHost, c.MongoPort, c.Database)
	adderService := adder.NewService(repo, logger)
	fetcherService := fetcher.NewService(repo, logger)
	modifierService := modifier.NewService(repo, logger)
	eraserService := eraser.NewService(repo, logger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		srv := grpc.NewServer(c.Port, adderService, fetcherService, modifierService, eraserService, logger)
		return srv.Serve()

	})

	logger.Fatal(g.Wait().Error())

}

func newMongo(host, port, database string) domain.Repository {
	db, cancel := mongo.NewConn(host, port, database)
	defer cancel()
	return mongo.NewMongo(db)
}
