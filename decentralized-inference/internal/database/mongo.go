package database

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type Database struct {
	mongo     *mongo.Client
	dbName    string
	dbAddress string
}

func (s *Database) GetMongo() *mongo.Client {
	return s.mongo
}

func (s *Database) SetMongo(mongo *mongo.Client) {
	s.mongo = mongo
}

func (s *Database) GetDB() *mongo.Database {
	return s.mongo.Database(s.dbName)
}

func InitMongo(dbname string, mongoAddr string) (*Database, error) {
	var s = &Database{
		dbName:    dbname,
		dbAddress: mongoAddr,
	}
	mongo, err := ConnectMongo(dbname, mongoAddr)
	if err != nil {
		return nil, err
	}
	s.SetMongo(mongo)
	return s, nil
}

func ConnectMongo(dbName string, mongoAddr string) (*mongo.Client, error) {
	wc := writeconcern.New(writeconcern.W(1), writeconcern.J(true))
	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(mongoAddr).SetWriteConcern(wc))
	if err != nil {
		return nil, err
	}
	_, client, _, err := mgm.DefaultConfigs()
	if err != nil {
		return nil, err
	}
	return client, nil
}
