package userservice

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"bookingrooms/internal/models"
	"bookingrooms/internal/repositories/mockrepository"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/mock"
)

func TestService_Register(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//prepare data
		register := models.Register{Username: "username", Password: "1234"}
		mockRepository := new(mockrepository.UserRepository)
		mockRepository.
			On("Register", mock.Anything).
			Return(nil)
		s := NewUserService(mockRepository)
		//actual
		actual := s.Register(register)
		//result
		assert.Equal(t, nil, actual)
		// เป็นการเช็คว่าใน Mock มีการส่งค่ามาถูกต้องหรือไม่ กรณีส่ง mock มาเกินจะเกิด error
		mockRepository.AssertExpectations(t)
	})
	t.Run("error", func(t *testing.T) {
		//prepare data
		register := models.Register{Username: "username", Password: "1234"}
		mockRepository := new(mockrepository.UserRepository)
		mockRepository.
			On("Register", mock.Anything).
			Return(errors.New("username is duplicated"))
		s := NewUserService(mockRepository)
		//actual
		actual := s.Register(register)
		//result
		assert.NotEqual(t, nil, actual)
		mockRepository.AssertExpectations(t)
	})
}

func TestService_Login(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		//assert
		expected := jwt.StandardClaims{
			ExpiresAt: 1671408000,
			IssuedAt:  1671148800,
			Subject:   "60",
		}
		//prepare data
		t.Setenv("TIME_TEST", fmt.Sprintf("%d-%d-%dT00:00:00Z", time.Now().Year(), time.Now().Month(), time.Now().Day()))
		defer os.Unsetenv("TIME_TEST")
		login := models.Login{Username: "username", Password: "1234"}
		mockRepository := new(mockrepository.UserRepository)
		mockRepository.On("GetPasswordHash", mock.Anything).
			Return(
				models.Users{
					Model: models.Model{
						ID:        60,
						CreatedAt: time.Date(2022, 11, 4, 15, 54, 56, 0, time.Local),
						UpdatedAt: time.Date(2022, 11, 8, 17, 14, 22, 0, time.Local),
					},
					FirstName:    "uracha",
					LastName:     "pudidvatanachok",
					NamePrefix:   "ms",
					SumGrade:     "A",
					AGrade:       "A",
					BGrade:       "A",
					CGrade:       "A",
					Name:         "a",
					Username:     "username",
					PasswordHash: "$2a$14$BU5cZk1dgGuXb0N0yAyvkehIqzQQmqyPWcVCcDArkiegj3WIdxVdy",
					CreatedBy:    60,
					UpdatedBy:    60,
				}, nil)
		mockRepository.On("SaveAccessToken", mock.Anything).
			Return(nil)
		mockRepository.On("SetRevokedTokenToRedis", mock.Anything, mock.Anything).
			Return(nil)
		s := NewUserService(mockRepository)
		//actual
		token, err := s.Login(login)
		actual := decodeTokenForTest(token)
		//result
		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
		mockRepository.AssertExpectations(t)
	})
	t.Run("error incorrect hash password", func(t *testing.T) {
		expected := ""
		expectedErr := "incorrect hash password"

		login := models.Login{Username: "username", Password: "1234"}
		mockRepository := new(mockrepository.UserRepository)
		mockRepository.On("GetPasswordHash", login.Username).
			Return(models.Users{PasswordHash: "notMatchPasswordHash"}, nil)
		s := NewUserService(mockRepository)

		actual, err := s.Login(login)

		assert.Equal(t, expectedErr, err.Error())
		assert.Equal(t, expected, actual)
		mockRepository.AssertExpectations(t)
	})
	t.Run("error Token is expired", func(t *testing.T) {
		expected := ""
		expectedErr := "Token is expired"

		t.Setenv("TIME_TEST", "2000-01-01T00:00:00Z")
		defer os.Unsetenv("TIME_TEST")
		login := models.Login{Username: "username", Password: "1234"}
		mockRepository := new(mockrepository.UserRepository)
		mockRepository.On("GetPasswordHash", login.Username).
			Return(models.Users{PasswordHash: "$2a$14$BU5cZk1dgGuXb0N0yAyvkehIqzQQmqyPWcVCcDArkiegj3WIdxVdy"}, nil)
		s := NewUserService(mockRepository)

		actual, err := s.Login(login)

		assert.Equal(t, expectedErr, err.Error())
		assert.Equal(t, expected, actual)
		mockRepository.AssertExpectations(t)
	})
}

func TestService_GetUser(t *testing.T) {
	//prepare data
	id := 60
	mockRepo := models.Users{
		Model: models.Model{
			ID:        60,
			CreatedAt: time.Date(2022, 11, 4, 15, 54, 56, 0, time.Local),
			UpdatedAt: time.Date(2022, 11, 8, 17, 14, 22, 0, time.Local),
		},
		FirstName:    "nabodee",
		LastName:     "srivichai",
		NamePrefix:   "mr",
		SumGrade:     "A",
		AGrade:       "A",
		BGrade:       "A",
		CGrade:       "A",
		Name:         "a",
		Username:     "moji1",
		PasswordHash: "$2a$14$BU5cZk1dgGuXb0N0yAyvkehIqzQQmqyPWcVCcDArkiegj3WIdxVdy",
		CreatedBy:    60,
		UpdatedBy:    60,
	}

	t.Run("success", func(t *testing.T) {
		//expected
		expected := models.Users{
			Model: models.Model{
				ID:        60,
				CreatedAt: time.Date(2022, 11, 4, 15, 54, 56, 0, time.Local),
				UpdatedAt: time.Date(2022, 11, 8, 17, 14, 22, 0, time.Local),
			},
			FirstName:    "nabodee",
			LastName:     "srivichai",
			NamePrefix:   "mr",
			SumGrade:     "A",
			AGrade:       "A",
			BGrade:       "A",
			CGrade:       "A",
			Name:         "a",
			Username:     "moji1",
			PasswordHash: "$2a$14$BU5cZk1dgGuXb0N0yAyvkehIqzQQmqyPWcVCcDArkiegj3WIdxVdy",
			CreatedBy:    60,
			UpdatedBy:    60,
		}

		//actual
		mockRepository := new(mockrepository.UserRepository)
		mockRepository.On("GetUser", id).Return(mockRepo, nil)
		s := NewUserService(mockRepository)
		actual, err := s.GetUser(id)

		//assert
		assert.Equal(t, nil, err)
		assert.Equal(t, expected, actual)
		mockRepository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		//expected
		expected := models.Users{}

		//actual
		mockRepository := new(mockrepository.UserRepository)
		mockRepository.On("GetUser", id).Return(models.Users{}, errors.New("record not found"))
		s := NewUserService(mockRepository)
		actual, err := s.GetUser(id)

		//assert
		assert.Equal(t, errors.New("record not found"), err)
		assert.Equal(t, expected, actual)
		mockRepository.AssertExpectations(t)
	})
}

func decodeTokenForTest(token string) jwt.StandardClaims {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		log.Println(err)
	}
	return jwt.StandardClaims{
		ExpiresAt: int64(claims["exp"].(float64)),
		Subject:   claims["sub"].(string),
		IssuedAt:  int64(claims["iat"].(float64)),
	}
}
