package model

type User struct {
	BaseModel
	Name     string `json:"name" gorm:"varchar(255);not null"`
	Email    string `json:"email" gorm:"varchar(255);not null;unique"`
	Password string `json:"password" gorm:"varchar(255);not null"`
	Role     string `json:"role" gorm:"enum('admin', 'user');default:'user'"`
}
