package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserId"`
	UserId    uint    `gorm:"not full"`
	Product   Product `gorm:"foreignKey:ProductId"`
	ProductId uint    `gorm:"not null"`
	Boss      User    `gorm:"ForeignKey:BossId"`
	BossId    uint    `gorm:"not null"`
}
