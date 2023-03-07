package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct{}

func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

func (p PostgresDB) DBConn() (*gorm.DB, error) {
	dsn := "host=localhost user=kai password=kai123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error when connect db: %v\n", err)
		return nil, err
	}

	fmt.Println("DB Connect Success!")

	return db, nil
}
