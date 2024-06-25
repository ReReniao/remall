package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId    uint `gorm:"not null"`
	AddressID uint `gorm:"not null"`
	Num       int
	OrderNum  uint64
	Type      uint // 已支付 1 未支付 2
	Money     float64
}
