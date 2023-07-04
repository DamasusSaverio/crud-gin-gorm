package models

import "gorm.io/gorm"

type Products struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	NamaProduct string `type:"varchar(255)" json:"nama_produk"`
	Deskripsi   string `type:"text" json:"deskripsi"`
	gorm.Model
}
