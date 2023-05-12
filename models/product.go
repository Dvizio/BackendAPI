package models

type Campaign struct {
	Id           int64  `gorm:"primaryKey" json:"id"`
	NamaCampaign string `gorm:"type:varchar(300)" json:"nama_campaign"`
	ClickThrough int64  `gorm:"type:int" json:"click_through"`
	Conversion   int64  `gorm:"type:int" json:"conversion"`
	NilaiAkhir   int64  `gorm:"type:int" json:"nilai_akhir"`
}
