package entity

type Candidate struct{
	ID  uint `json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	PartyID uint `json:"party_id"`
	Party Party  `gorm:"many2many" json:"party_name"`
}
