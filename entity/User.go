package entity

type User struct {
	Id       uint `gorm:"primary_key;auto_increment js" json:"id"`
	FullName string `gorm:"type:varchar(255);not null" json:"full_name"`
	Age      uint    `json:"age"`
	Email    string `gorm:"type:varchar(255);not null; unique" json:"email"`
	Phone    string `gorm:"type:varchar(255);not null; unique" json:"phone"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	RoleID   int    `json:"role_id"`
	Role     Role   `gorm:"many2many" json:"roles"`
}


