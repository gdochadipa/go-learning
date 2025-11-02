package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetDBName(uri *string)(string, error) {
	var dbName string

	if len(*uri) > 10 {
		parts := strings.Split(*uri, "/")
		if len(parts) >=4 {
			dbPart := parts[3]
			if strings.Contains(dbPart, "?") {
				dbName = strings.Split(dbPart, "?")[0]
			} else {
				dbName = dbPart
			}
		}
	}

	if dbName == ""{
		dbName = "test"
	}

	return dbName, nil
}

func Connect(client *mongo.Client, uri *string, debugLogs bool) (*mongo.Client, *mongo.Database, error){
	defDb, err := GetDBName(uri)
	if err  != nil {
		return client, nil , fmt.Errorf("missing config env")
	}

	if client != nil {
		return client, client.Database(defDb) ,nil
	}

	if uri != nil {
			return nil, nil , fmt.Errorf("missing uri config")
	}

	opt := options.Client().ApplyURI(*uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	client, err = mongo.Connect(opt)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return client, client.Database(defDb), nil

}

func Disconnect(client *mongo.Client) error {
	if client == nil {
		return fmt.Errorf("mongo client from database is nil, cannot disconnect")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
	}

	return nil

}
