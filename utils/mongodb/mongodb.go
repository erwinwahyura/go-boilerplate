package mongodb

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	// MongoConfig mongo config
	MongoConfig struct {
		Driver, Host, Port, Username, Password, DatabaseName string
	}
)

// NewMongoDB mongo db
func NewMongoDsn(dsn, dbname string) *mongo.Database {
	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal().Msgf("error when mongo.Connect(ctx, clientOptions), request: %v, error: %v", dsn, err.Error())
	}

	mongoDB := client.Database(dbname)

	log.Info().Msgf("ping mongodb %s", dbname)
	if err := mongoDB.Client().Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal().Msgf("error when ping mongodb %s, error: %v", dbname, err.Error())
	}

	return mongoDB
}
