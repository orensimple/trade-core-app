package model

import "github.com/google/uuid"

// Tabler is interface of GORM table name
type Tabler interface {
	TableName() string
}

// User is the model of User
type User struct {
	ID        uuid.UUID `gorm:"type:uuid"`
	Email     string    `gorm:"type:text;not null"`
	Password  string    `gorm:"type:text;not null"`
	FirstName string    `gorm:"type:text;not null"`
	LastName  string    `gorm:"type:text;not null"`
	Passport  int       `gorm:"type:int;not null"`
	Male      bool      `gorm:"type:bool"`
	About     string    `gorm:"type:text"`
	Address   string    `gorm:"type:text"`
}

// TableName gets table name of User
func (User) TableName() string {
	return "users"
}
