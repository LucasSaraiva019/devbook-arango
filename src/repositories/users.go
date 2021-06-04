package repositories

import (
	"errors"

	"github.com/LucasSaraiva019/devbook-arango/src/database"
	"github.com/LucasSaraiva019/devbook-arango/src/models"
)

func Create(user *models.User) error {
	tx := database.DB.Create(user)
	if tx == nil {
		return errors.New("Error creating User")
	}

	return tx.Error
}
