package seeder

import (
	"log"

	"github.com/ADMex1/GoProject/config"
	"github.com/ADMex1/GoProject/models"
	"github.com/ADMex1/GoProject/utils"
	"github.com/google/uuid"
)

func AdminSeeder() {
	password, _ := utils.HashPassword("useradmin12345")
	admin := models.User{
		Name:     "@ADMex1",
		Email:    "ADMex1@gmail.com",
		Password: password,
		Role:     "admin",
		PublicID: uuid.New(),
	}
	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("User seeded")
	}
}
