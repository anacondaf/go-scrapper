package db

import (
	"fmt"
	"github.com/kainguyen/go-scrapper/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	config *config.Config
}

func NewPostgresDB(config *config.Config) *PostgresDB {
	return &PostgresDB{config: config}
}

func (p PostgresDB) DBConn() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: p.config.DBUrl}), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error when connect db: %v\n", err)
		return nil, err
	}

	fmt.Println("DB Connect Success!")

	return db, nil
}
