package initializers

import "github.com/brenommelo/jwt-auth-go/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
