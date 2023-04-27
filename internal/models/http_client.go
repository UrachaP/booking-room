package models

type UsersHTTPClient struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status string `json:"status"`
}

type PostsHTTPClient struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type UserPosts struct {
	ID     int               `json:"id"`
	Name   string            `json:"name"`
	Email  string            `json:"email"`
	Gender string            `json:"gender"`
	Status string            `json:"status"`
	Posts  []PostsHTTPClient `json:"posts"`
}
