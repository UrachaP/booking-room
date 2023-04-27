package userservice

import (
	"os"
	"strconv"
	"time"

	"bookingrooms/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func (s service) createToken(id int) (string, error) {
	timeNow, err := now()
	if err != nil {
		return "", err
	}

	// Set custom claims
	claims := jwt.StandardClaims{
		ExpiresAt: timeNow.Add(time.Hour * 72).Unix(),
		Id:        uuid.NewString(),
		IssuedAt:  timeNow.Unix(),
		Subject:   strconv.Itoa(id),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s service) decodeToken(token string) (models.AccessTokens, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return models.AccessTokens{}, err
	}

	userId, err := strconv.Atoi(claims["sub"].(string))
	if err != nil {
		return models.AccessTokens{}, err
	}
	accessToken := models.AccessTokens{
		UserID:  userId,
		IDToken: claims["jti"].(string),
	}
	return accessToken, nil
}

func (s service) RevokedAccessToken(idToken string) error {
	return s.repository.RevokedAccessToken(idToken)
}

func now() (time.Time, error) {
	if os.Getenv("TIME_TEST") != "" {
		return time.Parse(time.RFC3339, os.Getenv("TIME_TEST"))
	}
	return time.Now(), nil
}
