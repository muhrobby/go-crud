package migration

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/muhrobby/go-crud/database"
	"github.com/muhrobby/go-crud/model"
)

func AutoMigrate() {
	err := database.ConnectDB().AutoMigrate(
		&model.User{},
	)

	if err != nil {
		log.Info("Failed to migrate")
	}

	log.Info("Successfully migrated")
}
