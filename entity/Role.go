package entity

type Role struct{
	ID uint `json:"id"`
	Name string `gorm:"varchar(255);not null" json:"name"`
}
