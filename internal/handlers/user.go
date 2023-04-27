package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"bookingrooms/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (h Handlers) Register(c echo.Context) error {
	var register models.Register
	err := c.Bind(&register)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(register)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.userService.Register(register)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "register success")
}

func (h Handlers) Login(c echo.Context) error {
	var login models.Login

	err := c.Bind(&login)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(login)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := h.userService.Login(login)
	if err == errors.New("incorrect hash password") {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func (h Handlers) GetLoginHistory(c echo.Context) error {
	loginHistory, err := h.userService.GetLoginHistory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &loginHistory)
}

func (h Handlers) RevokedAccessToken(c echo.Context) error {
	idToken := c.Param("id_token")
	err := h.userService.RevokedAccessToken(idToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "update revoked success")
}

func (h Handlers) GetUser(c echo.Context) error {
	requestId := c.Param("id")

	id, err := strconv.Atoi(requestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user, err := h.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h Handlers) GetUserList(c echo.Context) error {
	var pagination models.Pagination
	err := c.Bind(&pagination)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	users, err := h.userService.GetUsers(pagination)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, &users)
}

func (h Handlers) UpdateUser(c echo.Context) error {
	var user models.Users
	// read param
	requestId := c.Param("id")

	// convert string to int
	id, err := strconv.Atoi(requestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user.ID = id

	// check duplicate
	count, err := h.userService.GetCountUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if count == 0 {
		return c.JSON(http.StatusNotFound, errors.New("no data").Error())
	}

	// read request body and use bind can validate type value
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// validate required
	err = c.Validate(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// get token
	userClaims := c.Get("user").(*jwt.Token)
	claims := userClaims.Claims.(*jwt.StandardClaims)
	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	user.UpdatedBy = userId

	// check have image_id in db
	//err = h.userService.TempImageRepository.HaveTempImages(user.ImageID)
	err = h.tempImageService.HaveTempImages(user.ImageID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// update table user
	err = h.userService.UpdateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	tempImage := h.tempImageService.NewTempImage(user.ImageID)

	err = h.tempImageService.UpdateTempImages(tempImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "update user success")
}

func (h Handlers) UpdateUserGrade(c echo.Context) error {
	var userGrade models.UserGrade
	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = c.Bind(&userGrade)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(userGrade)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwt.StandardClaims)
	userId, err := strconv.Atoi(claims.Subject)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userGrade.ID = id
	userGrade.UpdatedBy = userId

	err = h.userService.UpdateUserGrade(userGrade)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "update user success")
}

func (h Handlers) DeleteUser(c echo.Context) error {
	requestId := c.Param("id")
	id, err := strconv.Atoi(requestId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = h.userService.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.String(http.StatusOK, "deleted user success")
}

func (h Handlers) PreloadUserBookings(c echo.Context) error {
	userBooking, err := h.userService.PreloadUserBookings()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userBooking)
}

func (h Handlers) CreateUserWithBooking(c echo.Context) error {
	var userWithBooking models.UserWithBookings
	err := c.Bind(&userWithBooking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(userWithBooking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	register := models.Register{Username: userWithBooking.Username, Password: userWithBooking.Password}
	err = h.userService.Register(register)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userId, err := h.userService.GetUserIdByUsername(register.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.bookingService.CreateBookingWithUserId(userWithBooking, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "register and save booking success")
}

func (h Handlers) GetRevokedTokenFromRedis(idToken string) (string, error) {
	return h.userService.GetRevokedTokenFromRedis(idToken)
}
