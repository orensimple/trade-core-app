package domain

import "github.com/google/uuid"

// User is user in app.
type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	FirstName string
	LastName  string
	Male      bool
	About     string
	Address   string
}

// LoginRequest struct for login.
type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// UserSearchRequest struct for full text users search.
type UserSearchRequest struct {
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
}
