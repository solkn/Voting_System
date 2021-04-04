package entity

type Result struct{
	ID uint `json:"id"`
	Ballot uint `json:"ballot"`
	PartyID uint `json:"party_id"`
	Party Party  `gorm:"many2many" json:"party_name"`

}