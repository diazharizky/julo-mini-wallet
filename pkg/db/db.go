package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbClient struct {
	host     string
	port     int
	user     string
	password string
	dbName   string
	sslMode  string
}

func New() dbClient {
	return dbClient{
		host:     "localhost",
		port:     5432,
		user:     "julo",
		password: "julo",
		dbName:   "julo",
		sslMode:  "disable",
	}
}

func (client dbClient) DB() *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(client.dsn()), &gorm.Config{},
	)
	if err != nil {
		log.Fatalf("error unable to initialize db: %v", err)
	}

	return db
}

func (client dbClient) dsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		client.host,
		client.port,
		client.user,
		client.password,
		client.dbName,
		client.sslMode,
	)
}
