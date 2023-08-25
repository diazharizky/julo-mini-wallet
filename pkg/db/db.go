package db

import (
	"fmt"
	"log"

	"github.com/diazharizky/julo-mini-wallet/config"
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

func init() {
	config.Global.SetDefault("postgres.sslMode", "disable")
}

func New() dbClient {
	return dbClient{
		host:     config.Global.GetString("postgres.host"),
		port:     config.Global.GetInt("postgres.port"),
		user:     config.Global.GetString("postgres.user"),
		password: config.Global.GetString("postgres.password"),
		dbName:   config.Global.GetString("postgres.db"),
		sslMode:  config.Global.GetString("postgres.sslMode"),
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
