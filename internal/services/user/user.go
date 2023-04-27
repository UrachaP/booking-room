package userservice

import (
	"errors"
	"log"

	"bookingrooms/internal/models"
	userrepository "bookingrooms/internal/repositories/user"
	"github.com/go-redis/redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository userrepository.UserRepository
}

func NewUserService(repository userrepository.UserRepository) service {
	return service{repository: repository}
}

type UserService interface {
	Register(register models.Register) error
	Login(login models.Login) (string, error)
	GetUser(id int) (models.Users, error)
	NewUser(username string, passwordHash string) *models.Users
	GetUserIdByUsername(username string) (int, error)
	GetLoginHistory() (*[]models.LoginHistory, error)
	RevokedAccessToken(idToken string) error
	GetUsers(pagination models.Pagination) (*[]models.Users, error)
	GetCountUserById(id int) (int, error)
	UpdateUser(user models.Users) error
	GetRevokedTokenFromRedis(idToken string) (string, error)
	SetRevokedTokensToRedis(accessTokens *[]models.LoginHistory) error
	GetSumGrade(g models.UserGrade) string
	UpdateUserGrade(userGrade models.UserGrade) error
	DeleteUser(id int) error
	PreloadUserBookings() ([]models.UserBookings, error)
}

func (s service) Register(register models.Register) error {
	// hash password and new user
	passwordHash, err := s.hashPassword(register.Password)
	if err != nil {
		return err
	}
	user := s.NewUser(register.Username, passwordHash)
	return s.repository.Register(*user)
}

func (s service) Login(login models.Login) (string, error) {
	user, err := s.repository.GetPasswordHash(login.Username)
	if err != nil {
		return "", err
	}

	if !s.checkPasswordHash(login.Password, user.PasswordHash) {
		return "", errors.New("incorrect hash password")
	}

	token, err := s.createToken(user.ID)
	if err != nil {
		return "", err
	}
	accessToken, err := s.decodeToken(token)
	if err != nil {
		return "", err
	}

	err = s.repository.SaveAccessToken(accessToken)
	if err != nil {
		return "", err
	}

	err = s.repository.SetRevokedTokenToRedis(accessToken.IDToken, "0")
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s service) GetUser(id int) (models.Users, error) {
	return s.repository.GetUser(id)
}

func (s service) NewUser(username string, passwordHash string) *models.Users {
	return &models.Users{
		Username:     username,
		PasswordHash: passwordHash,
	}
}

func (s service) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s service) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s service) GetUserIdByUsername(username string) (int, error) {
	user, err := s.repository.GetUserByUsername(username)
	return user.ID, err
}

func (s service) GetLoginHistory() (*[]models.LoginHistory, error) {
	return s.repository.GetLoginHistory()
}

func (s service) GetUsers(pagination models.Pagination) (*[]models.Users, error) {
	return s.repository.GetUserList(pagination)
}

func (s service) GetCountUserById(id int) (int, error) {
	return s.repository.GetCountUserById(id)
}

func (s service) UpdateUser(user models.Users) error {
	return s.repository.UpdateUser(user)
}

func (s service) SetRevokedTokensToRedis(accessTokens *[]models.LoginHistory) error {
	var err error
	for _, history := range *accessTokens {
		err = s.repository.SetRevokedTokenToRedis(history.IDToken, history.Revoked)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s service) GetRevokedTokenFromRedis(idToken string) (string, error) {
	revoked, err := s.repository.GetRevokedTokenFromRedis(idToken)
	switch { //should use if
	case err == redis.Nil:
		log.Println("key does not exist")
		//get revoked to set redis
		revoked, err = s.repository.GetRevokedToken(idToken)
		if err != nil {
			return "", err
		}
		return revoked, s.repository.SetRevokedTokenToRedis(idToken, revoked)
	case err != nil:
		log.Println("Get failed", err)
		return "", err
		//case revoked == "":
		//	log.Println("value is empty")
		//	return "", err
	}
	return revoked, nil
}

func (s service) UpdateUserGrade(userGrade models.UserGrade) error {
	count, err := s.repository.GetCountUserById(userGrade.ID)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("no data")
	}
	userGrade.SumGrade = s.GetSumGrade(userGrade)
	return s.repository.SaveUserGrade(userGrade)
}

func (s service) DeleteUser(id int) error {
	return s.repository.DeleteUserByID(id)
}

func (s service) PreloadUserBookings() ([]models.UserBookings, error) {
	return s.repository.PreloadUserBookings()
}
