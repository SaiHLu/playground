package database

import (
	"fmt"
	"log"

	"github.com/playground/fiber/common/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	*gorm.DB
}

func SetupDB() {
	env := config.GetEnv().DBConfig
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		env.Host,
		env.Port,
		env.User,
		env.Password,
		env.DbName,
		env.SSLMode,
	)

	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Could not connect to database.")
	}
}
