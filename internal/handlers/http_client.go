package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"bookingrooms/internal/models"
	"github.com/labstack/echo/v4"
)

func (h Handlers) GetUserHTTPClient(c echo.Context) error {
	response, err := http.Get("https://gorest.co.in/public/v2/users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)

	var users []models.UsersHTTPClient

	//err = json.NewDecoder(response.Body).Decode(&users)
	//if err != nil {
	//	return c.JSON(http.StatusInternalServerError, err.Error())
	//}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = json.Unmarshal(body, &users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (h Handlers) GetPostHTTPClient(c echo.Context) error {
	response, err := http.Get("https://gorest.co.in/public/v2/posts")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer response.Body.Close()
	var post []models.PostsHTTPClient

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	err = json.Unmarshal(body, &post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	log.Println(post)
	return c.JSON(http.StatusOK, post)
}

func (h Handlers) GetUserPostsHTTPClient(c echo.Context) error {
	//get users
	body, err := DecodeResponseBody("https://gorest.co.in/public/v2/users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var userPosts []models.UserPosts
	err = json.Unmarshal(body, &userPosts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	//get posts
	body, err = DecodeResponseBody("https://gorest.co.in/public/v2/posts")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	var posts []models.PostsHTTPClient
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for _, post := range posts {
		for i := 0; i < len(userPosts); i++ {
			if post.UserID == userPosts[i].ID {
				userPosts[i].Posts = append(userPosts[i].Posts, post)
			}
		}
	}

	return c.JSON(http.StatusOK, userPosts)
}

func DecodeResponseBody(html string) ([]byte, error) {
	response, err := http.Get(html)
	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
