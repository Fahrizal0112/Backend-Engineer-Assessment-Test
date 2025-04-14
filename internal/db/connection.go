package db

import (
	"banking-service/internal/config"
	"banking-service/internal/models"
	"banking-service/pkg/logger"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connect(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
	)
	logger.Info("Connection to database", "dsn", fmt.Sprintf("host=%s port=%s user=%s dbname=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Name,
		))

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})

		if err != nil {
			logger.Critical(fmt.Sprintf("Failed to connect to database: %v", err))
			return nil, err
		}
		logger.Info("Database connection established")

		if err := db.AutoMigrate(&models.Nasabah{}); err != nil {
			logger.Critical(fmt.Sprintf("Failed to migrate database: %v", err))
			return nil, err
		}
		logger.Info("Database migration completed")
		return db, nil
}
