package migrations

import (
	"github.com/st4rgaze/otaqu/app/models"
	"github.com/st4rgaze/otaqu/config"
)

func CreateHotelsTable() error {
	db := config.DB
	return db.AutoMigrate(&models.Hotel{})
}
