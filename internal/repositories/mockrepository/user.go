package mockrepository

import (
	"bookingrooms/internal/models"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (mock *UserRepository) Register(user models.Users) error {
	argument := mock.Called(user)
	return argument.Error(0)
}
func (mock *UserRepository) GetPasswordHash(username string) (models.Users, error) {
	argument := mock.Called(username)
	return argument.Get(0).(models.Users), argument.Error(1)
}
func (mock *UserRepository) SaveAccessToken(accessToken models.AccessTokens) error {
	argument := mock.Called(accessToken)
	return argument.Error(0)
}
func (mock *UserRepository) GetLoginHistory() (*[]models.LoginHistory, error) {
	argument := mock.Called()
	return argument.Get(0).(*[]models.LoginHistory), argument.Error(1)
}
func (mock *UserRepository) RevokedAccessToken(idToken string) error {
	argument := mock.Called(idToken)
	return argument.Error(0)
}
func (mock *UserRepository) GetUserList(pagination models.Pagination) (*[]models.Users, error) {
	argument := mock.Called(pagination)
	return argument.Get(0).(*[]models.Users), argument.Error(1)
}
func (mock *UserRepository) GetCountUserById(id int) (int, error) {
	argument := mock.Called(id)
	return argument.Get(0).(int), argument.Error(1).(error)
}
func (mock *UserRepository) GetUser(id int) (models.Users, error) {
	argument := mock.Called(id)
	return argument.Get(0).(models.Users), argument.Error(1)
}
func (mock *UserRepository) GetRevokedToken(id string) (string, error) {
	argument := mock.Called(id)
	return argument.Get(0).(string), argument.Error(1)
}
func (mock *UserRepository) UpdateUser(user models.Users) error {
	argument := mock.Called(user)
	return argument.Error(0)
}
func (mock *UserRepository) SaveUserGrade(userGrade models.UserGrade) error {
	argument := mock.Called(userGrade)
	return argument.Error(0)
}
func (mock *UserRepository) DeleteUserByID(id int) error {
	argument := mock.Called(id)
	return argument.Error(0)
}
func (mock *UserRepository) GetUserByUsername(username string) (models.Users, error) {
	argument := mock.Called(username)
	return argument.Get(0).(models.Users), argument.Error(1)
}
func (mock *UserRepository) PreloadUserBookings() ([]models.UserBookings, error) {
	argument := mock.Called()
	return argument.Get(0).([]models.UserBookings), argument.Error(1)
}
func (mock *UserRepository) SetRevokedTokenToRedis(idToken, revoked string) error {
	argument := mock.Called(idToken, revoked)
	return argument.Error(0)
}
func (mock *UserRepository) GetRevokedTokenFromRedis(idToken string) (string, error) {
	argument := mock.Called(idToken)
	return argument.Get(0).(string), argument.Error(1)
}
func (mock *UserRepository) GetUserById(id int) (models.Users, error) {
	argument := mock.Called(id)
	return argument.Get(0).(models.Users), argument.Error(1)
}
