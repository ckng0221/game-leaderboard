package initializers

import "api/models"

func SynDatabase() {
	Db.AutoMigrate(
		&models.User{},
		&models.Gameplay{},
		&models.Score{},
	)
}
