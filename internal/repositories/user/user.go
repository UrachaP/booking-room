package userrepository

import (
	"errors"

	"bookingrooms/internal/models"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUserRepository(db *gorm.DB, rdb *redis.Client) UserRepository {
	return repository{db: db, rdb: rdb}
}

type UserRepository interface {
	Register(user models.Users) error
	GetPasswordHash(username string) (models.Users, error)
	SaveAccessToken(accessToken models.AccessTokens) error
	GetLoginHistory() (*[]models.LoginHistory, error)
	RevokedAccessToken(idToken string) error
	GetUserList(pagination models.Pagination) (*[]models.Users, error)
	GetCountUserById(id int) (int, error)
	GetUser(id int) (models.Users, error)
	GetRevokedToken(id string) (string, error)
	UpdateUser(user models.Users) error
	SaveUserGrade(userGrade models.UserGrade) error
	DeleteUserByID(id int) error
	GetUserByUsername(username string) (models.Users, error)
	PreloadUserBookings() ([]models.UserBookings, error)
	SetRevokedTokenToRedis(idToken, revoked string) error
	GetRevokedTokenFromRedis(idToken string) (string, error)
}

func (r repository) GetUser(id int) (models.Users, error) {
	var user models.Users
	return user, r.db.First(&user, id).Error
}

func (r repository) Register(user models.Users) error {
	//start transaction
	return r.db.Transaction(func(tx *gorm.DB) error {
		var count int
		err := tx.Model(&models.Users{}).Select("COUNT(*)").Where("username = ?", user.Username).First(&count).Error
		if err != nil {
			// return any error will roll back
			return err
		}
		if count != 0 {
			return errors.New("username is duplicated")
		}

		// created user
		err = tx.Create(&user).Error
		if err != nil {
			return err
		}

		// updated user created_by
		err = tx.Model(&user).Update("created_by", user.ID).Error
		if err != nil {
			return err
		}
		// return nil will commit the whole transaction
		return nil
	})
}

func (r repository) GetCountUserById(id int) (int, error) {
	var count int
	return count, r.db.
		Model(&models.Users{}).
		Select("COUNT(*)").
		Where("id = ?", id).
		First(&count).Error
}

func (r repository) GetUserById(id int) (models.Users, error) {
	var user models.Users
	return user, r.db.First(&user, id).Error
}

func (r repository) GetUserList(pagination models.Pagination) (*[]models.Users, error) {
	var users *[]models.Users
	offset := pagination.Limit * (pagination.Page - 1)
	return users, r.db.
		Limit(pagination.Limit).
		Offset(offset).
		Order(pagination.Sort).
		Find(&users).
		Error
}

func (r repository) UpdateUser(user models.Users) error {
	return r.db.Updates(&user).Error
}

func (r repository) SaveUserGrade(userGrade models.UserGrade) error {
	return r.db.Save(&userGrade).Error
}

func (r repository) DeleteUserByID(id int) error {
	return r.db.Delete(&models.Users{}, id).Error
}

func (r repository) SaveAccessToken(accessToken models.AccessTokens) error {
	return r.db.Create(&accessToken).Error
}

func (r repository) GetPasswordHash(username string) (models.Users, error) {
	var user models.Users
	err := r.db.
		Select("id", "password_hash").
		Where("username = ?", username).
		First(&user).Error
	return user, err
}

func (r repository) PreloadUserBookings() ([]models.UserBookings, error) {
	var userBookings []models.UserBookings
	return userBookings, r.db.Debug().
		Model(&models.Users{}).
		Select("id", "CONCAT(COALESCE(name_prefix,'') , first_name , ' ' , last_name) AS full_name").
		Preload("Bookings", func(db *gorm.DB) *gorm.DB {
			return db.Select("users_id", "CONCAT(start_date,' - ',end_date) AS booking_time", "room_name", "amount_person").
				Joins("INNER JOIN rooms ON rooms.id = rooms_id")
		}).
		Find(&userBookings).Error
}

func (r repository) GetLoginHistory() (*[]models.LoginHistory, error) {
	var loginHistory *[]models.LoginHistory
	return loginHistory, r.db.
		Select([]string{
			"CONCAT(COALESCE(name_prefix,''),first_name,last_name) AS full_name",
			"id_token",
			"access_tokens.created_at",
			"IF(revoked = '0','ใช้งาน','หยุดใช้งาน') AS revoked",
		}).
		Joins("LEFT JOIN exam.users ON exam.users.id = exam.access_tokens.user_id").
		Find(loginHistory).Error
}

func (r repository) GetUserByUsername(username string) (models.Users, error) {
	var user models.Users
	return user, r.db.Where("username = ?", username).First(&user).Error
}
