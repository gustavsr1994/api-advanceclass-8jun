package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NameProduct string `gorm:"type:varchar(20)" json:"name_product"`
	Qty         int    `gorm:"type:int" json:"qty"`
	DescProduct string `gorm:"type:text" json:"desc_product"`
}
