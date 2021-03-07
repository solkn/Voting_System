package entity

type Candidate struct{
	ID  uint `json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	PartyID uint `gorm:"not null;auto_preload" json:"partyID"`
	Party Party  `gorm:"auto_preload" json:"party"`
}
