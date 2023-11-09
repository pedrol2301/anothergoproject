package initializers

import "gotodolist/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.ToDo{})
}
