package db

import (
	"fmt"

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

func (client dbClient) DB() (*gorm.DB, error) {
	return gorm.Open(
		postgres.Open(client.dsn()), &gorm.Config{},
	)
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
