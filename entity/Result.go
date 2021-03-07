package entity

type Result struct{
	ID uint `json:"id"`
	Ballot uint `json:"vote"`
	PartyID uint `gorm:"not null;auto_preload" json:"party_id"`
	Party Party `gorm:"auto_preload" json:"party"`

}