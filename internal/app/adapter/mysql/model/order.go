package model

// Order is the model of orders
type Order struct {
	ID      string  `gorm:"column:order_id;type:uuid"`
	UserID  string  `gorm:"type:uuid"`
	User    User    // `gorm:"foreignKey:UserID;references:ID"`
	Payment Payment // `gorm:"foreignKey:OrderID;references:ID"`
}
