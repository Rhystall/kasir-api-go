package service_test

import (
	"api-kasirapp/input"
	"api-kasirapp/models"
	"api-kasirapp/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock repository
type mockUserRepository struct {
	findByEmailFunc func(string) (models.User, error)
}

func (m *mockUserRepository) FindByEmail(email string) (models.User, error) {
	return m.findByEmailFunc(email)
}

func (m *mockUserRepository) Save(user models.User) (models.User, error) {
	return models.User{}, nil
}

func (m *mockUserRepository) FindByID(id int) (models.User, error) {
	return models.User{}, nil
}

func (m *mockUserRepository) FindAll() ([]models.User, error) {
	return nil, nil
}

func (m *mockUserRepository) ActivateUser(ID int) (models.User, error) {
	return models.User{}, nil
}

func (m *mockUserRepository) FindByPhone(phone string) (models.User, error) {
	return models.User{}, nil
}


func TestIsEmailAvailable(t *testing.T) {
	t.Run("Email Available", func(t *testing.T) {
		mockRepo := &mockUserRepository{
			findByEmailFunc: func(email string) (models.User, error) {
				return models.User{}, nil // No user found
			},
		}

		userService := service.NewService(mockRepo)
		input := input.CheckEmailInput{Email: "available@example.com"}

		result, err := userService.IsEmailAvailable(input)

		assert.Nil(t, err)
		assert.True(t, result)
	})

	t.Run("Email Not Available", func(t *testing.T) {
		mockRepo := &mockUserRepository{
			findByEmailFunc: func(email string) (models.User, error) {
				return models.User{ID: 1, Email: email}, nil // User exists
			},
		}

		userService := service.NewService(mockRepo)
		input := input.CheckEmailInput{Email: "taken@example.com"}

		result, err := userService.IsEmailAvailable(input)

		assert.Nil(t, err)
		assert.False(t, result)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo := &mockUserRepository{
			findByEmailFunc: func(email string) (models.User, error) {
				return models.User{}, errors.New("db error")
			},
		}

		userService := service.NewService(mockRepo)
		input := input.CheckEmailInput{Email: "error@example.com"}

		result, err := userService.IsEmailAvailable(input)

		assert.NotNil(t, err)
		assert.False(t, result)
	})
}
