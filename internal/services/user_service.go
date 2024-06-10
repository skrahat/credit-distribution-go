package services

import (
	"credit-distribution-go/internal/models"
	"credit-distribution-go/internal/repositories"
)

func GetAllUsers() ([]models.User, error) {
    return repositories.GetAllUsers()
}

func GetUser(id string) (models.User, error) {
    return repositories.GetUser(id)
}
