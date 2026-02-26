package repository

import (
	"vita-track-ai/database"
	"vita-track-ai/models"
	"vita-track-ai/utility"
)

func GetUserModelByEmail(email string) (*models.User, error) {
	var user models.User
	tx := database.DB.Where("email = ?", email).First(&user)
	err := tx.Error

	return &user, err

}

func SaveUser(u *models.User) error {
	var err error

	if u.Password != nil {
		hashedPassword, err := utility.HashPassword(*u.Password)
		if err != nil {
			return err
		}

		u.Password = hashedPassword
	}
	err = database.DB.Create(u).Error
	return err
}
