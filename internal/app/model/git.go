package model

type Git struct {
	BaseModel
	Name  string `json:"name" gorm:"varchar(255);not null"`
	Url   string `json:"url" gorm:"varchar(255);not null"`
	Type  string `json:"type" gorm:"enum('github', 'gitlab', 'gitee', 'bitbucket');default:'github'"`
	Token string `json:"token" gorm:"varchar(255);not null"`
}
