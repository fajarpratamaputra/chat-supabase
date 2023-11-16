package postgresql

import (
	"chitchat/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	_ "github.com/lib/pq"
)

type Client struct {
	*gorm.DB
}

func NewClient() *Client {
	// Replace these with your PostgreSQL credentials
	dbUser := config.Cfg.GetString("PG_USER")
	dbPassword := config.Cfg.GetString("PG_PASS")
	dbName := config.Cfg.GetString("PG_DBNAME")
	dbHost := config.Cfg.GetString("PG_HOST") // Or your PostgreSQL host
	dbPort := config.Cfg.GetString("PG_PORT") // Or your PostgreSQL port

	// Create the connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}

	return &Client{DB: db}
}
