package seed

import (
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/utils"
	"log"

	"gorm.io/gorm"
)

func UserSeed(db *gorm.DB) {
	password, _ := utils.HashingPassword("password123")

	users := []entity.User{
		{
			Name:     "Admin User",
			Email:    "admin@example.com",
			Password: password,
			Address:  "Jl. Admin No. 1",
			Phone:    "081234567890",
		},
		{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: password,
			Address:  "Jl. John No. 2",
			Phone:    "081234567891",
		},
	}

	for _, user := range users {
		var existingUser entity.User
		err := db.Where("email = ?", user.Email).First(&existingUser).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				if err := db.Create(&user).Error; err != nil {
					log.Printf("Could not seed user %s: %v", user.Email, err)
				}
			} else {
				log.Printf("Error checking user %s: %v", user.Email, err)
			}
		}
	}
}
