package database

import (
	"github.com/erwinwahyura/go-boilerplate/app/model"
	"github.com/erwinwahyura/go-boilerplate/utils/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoCollection
type MongoCollection struct {
	MessageMaster *mongo.Database
	MessageSlave  *mongo.Database
}

// NewMongoCollection ...
func NewMongoCollection(config model.Config) MongoCollection {

	// DB
	master := mongodb.NewMongoDsn(config.Database.LogDB.Master, "dbname")
	slave := mongodb.NewMongoDsn(config.Database.LogDB.Slave, "dbname")

	return MongoCollection{
		MessageMaster: master,
		MessageSlave:  slave,
	}
}
