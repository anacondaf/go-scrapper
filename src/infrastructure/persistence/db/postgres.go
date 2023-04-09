package db

import (
	"fmt"
	"github.com/kainguyen/go-scrapper/src/config"
	"github.com/kainguyen/go-scrapper/src/core/domain/models"
	"github.com/kainguyen/go-scrapper/src/utils"
	"github.com/qustavo/dotsql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	config *config.Config
}

func NewPostgresDB(config *config.Config) *PostgresDB {
	return &PostgresDB{config: config}
}

func (p PostgresDB) LoadDefaultSQLCmd(db *gorm.DB) error {
	// Generic database interface
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	var sqlFilePath = fmt.Sprintf("%v/src/core/domain/commands/command.sql", utils.GetWorkDirectory())

	dot, err := dotsql.LoadFromFile(sqlFilePath)
	if err != nil {
		return err
	}

	// uuid-generate-extension is name of sql command
	_, err = dot.Exec(sqlDB, "uuid-generate-extension")
	if err != nil {
		return err
	}

	_, err = dot.Exec(sqlDB, "create-database")
	if err != nil {
		return err
	}

	return nil
}

func (p PostgresDB) DBConn() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: p.config.Database.ConnectionString}), &gorm.Config{
		NamingStrategy: p.config.Database.GormConfig.NamingStrategy,
	})

	err = p.LoadDefaultSQLCmd(db)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Post{}, &models.PostImages{})
	if err != nil {
		return nil, err
	}

	if err != nil {
		fmt.Printf("Error when connect db: %v\n", err)
		return nil, err
	}

	fmt.Println("DB Connect Success!")

	return db, nil
}
