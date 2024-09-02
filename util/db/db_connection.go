package util

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"

	"golang-template-service/config"
)

type DatabaseConnection struct {
	DbConnection *gorm.DB
}

// New constructs new DatabaseConnection
func New(config config.DatabaseConfig) (*DatabaseConnection, error) {
	connStr := Connect(config)

	return &DatabaseConnection{
		DbConnection: connStr,
	}, nil
}

func GetPGSQLConfig(config config.DatabaseConfig) string {
	pgsql := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DbName,
	)
	return pgsql
}

func Connect(config config.DatabaseConfig) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}
	if db, err = gorm.Open(postgres.Open(GetPGSQLConfig(config)), gormConfig); err != nil {

		log.Fatal(err)
	}
	return db
}
