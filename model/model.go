package model

import (
	"gorm.io/gorm"
)

// Product is a model of a product
type Product struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(100);not null"`
	Observations *string `gorm:"type:varchar(100)"`
	Price        int     `gorm:"not null"`
	InvoiceItems []InvoiceItem
}

// InvoiceHeader is a model of an invoice header
type InvoiceHeader struct {
	gorm.Model
	Client       string `gorm:"type:varchar(100);not null"`
	InvoiceItems []InvoiceItem
}

// InvoiceItem is a model of an invoice item
type InvoiceItem struct {
	gorm.Model
	InvoiceHeaderID uint
	ProductID       uint
}
