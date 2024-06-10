package repositories

import (
	"credit-distribution-go/internal/models"
	"errors"
)

var mockUsers = []models.User{
    {ID: "1", Name: "John Doe", Age: 30},
    {ID: "2", Name: "Jane Doe", Age: 25},
}

func GetAllUsers() ([]models.User, error) {
    return mockUsers, nil
}

func GetUser(id string) (models.User, error) {
    for _, user := range mockUsers {
        if user.ID == id {
            return user, nil
        }
    }
    return models.User{}, errors.New("user not found")
}
