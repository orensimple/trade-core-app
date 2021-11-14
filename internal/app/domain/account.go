package domain

import "github.com/google/uuid"

// Account is the model of Account
type Account struct {
	ID           uuid.UUID `gorm:"type:uuid;pk"`
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	AccountID    uuid.UUID `gorm:"type:uuid;not null"`
	CurrencyCode string    `gorm:"type:text"`

	User *User `json:"-" gorm:"foreignKey:UserID"`
}

// TableName gets table name of Account
func (Account) TableName() string {
	return "accounts"
}
