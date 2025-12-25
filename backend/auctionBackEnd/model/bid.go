package model

import (
	"time"

	"gorm.io/gorm"
)

type Bid struct {
	ID     uint       `json:"id" gorm:"primaryKey"`
	Buyer  string     `json:"buyer" gorm:"not null;size:42;index"`
	Bid    string     `json:"bid" gorm:"not null;size:78"`
	Deal   bool       `json:"deal" gorm:"not null;default:0"`
	DealAt *time.Time `json:"dealAt"`

	CreatedAt time.Time      `json:"createAt"`
	UpdatedAt time.Time      `json:"updateAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
