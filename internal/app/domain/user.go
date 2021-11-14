package domain

import "github.com/google/uuid"

// User is the model of User
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;pk"`
	Email     string    `gorm:"type:text;not null"`
	Password  string    `gorm:"type:text;not null"`
	Phone     string    `gorm:"type:text"`
	FirstName string    `gorm:"type:text;not null"`
	LastName  string    `gorm:"type:text;not null"`
	Passport  int       `gorm:"type:int;not null"`
	Male      bool      `gorm:"type:bool"`
	About     string    `gorm:"type:text"`
	Address   string    `gorm:"type:text"`

	Accounts []*Account `gorm:"foreignKey:UserID;association_foreignKey:ID"`
}

// TableName gets table name of User
func (User) TableName() string {
	return "users"
}

// RegisterRequest struct for registration.
type RegisterRequest struct {
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Phone     string `form:"phone" json:"phone"`
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Passport  int    `form:"passport" json:"passport" binding:"required"`
	Male      string `form:"male" json:"male"`
	About     string `form:"about" json:"about"`
	Address   string `form:"address" json:"address" binding:"required"`
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
