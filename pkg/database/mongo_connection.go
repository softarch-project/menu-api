package database

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/softarch-project/menu-api/config"
	"github.com/softarch-project/menu-api/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB(config *config.Config) (*mongo.Client, error) {
	connectionURI := util.NewConnectionUrlBuilder(config.Database)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to mongo database successfully")
	return client, nil
}
