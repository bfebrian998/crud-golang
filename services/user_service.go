package services

import (
	"go-crud/models"
	"go-crud/repository"
)

func CreateUser(user models.User) (models.User, error) {
	return repository.CreateUser(user)
}

func GetAllUsers() ([]models.User, error) {
	return repository.GetAllUsers()
}

func GetUserByID(id int) (models.User, error) {
	return repository.GetUserByID(id)
}

func UpdateUser(id int, user models.User) (models.User, error) {
	return repository.UpdateUser(id, user)
}

func DeleteUser(id int) error {
	return repository.DeleteUser(id)
}
