package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Products *mongo.Collection
var Users *mongo.Collection
var Orders *mongo.Collection

var MongoClient *mongo.Client

type TCollection struct {
	QrCodes *mongo.Collection
	Users   *mongo.Collection
}

type MongoConnection struct {
	Username string
	Password string
	Database string
	Host     string
	Srv      string
}

var Collections TCollection

func createConnection() (MongoConnection, error) {
	var mongoConnection MongoConnection
	var username = os.Getenv("USERNAME")
	var password = os.Getenv("PASSWORD")
	var database = os.Getenv("DATABASE")
	var host = os.Getenv("HOST")
	var credentials = username + ":" + password
	var errorTemplate = `"%s must be present in .env"`

	if len(username) == 0 {
		return mongoConnection, fmt.Errorf(errorTemplate, "USERNAME")
	}

	if len(password) == 0 {
		return mongoConnection, fmt.Errorf(errorTemplate, "PASSWORD")
	}

	if len(database) == 0 {
		return mongoConnection, fmt.Errorf(errorTemplate, "DATABASE")
	}

	if len(host) == 0 {
		return mongoConnection, fmt.Errorf(errorTemplate, "HOST")
	}

	return MongoConnection{
		Username: username,
		Password: password,
		Database: database,
		Host:     host,
		Srv:      fmt.Sprintf("mongodb+srv://%s@%s/%s", credentials, host, database),
	}, nil
}

func Connect() error {
	var err error
	var connection MongoConnection
	var clientOptions *options.ClientOptions

	connection, err = createConnection()
	if err != nil {
		return err
	}

	clientOptions = options.Client().ApplyURI(connection.Srv)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return err
	}

	MongoClient = client
	Collections.QrCodes = client.Database(connection.Database).Collection("qrcodes")
	Collections.Users = client.Database(connection.Database).Collection("users")

	return nil
}
