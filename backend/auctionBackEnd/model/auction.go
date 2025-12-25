package model

import (
	"time"

	"gorm.io/gorm"
)

type Auction struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	SellerAddress   string    `json:"sellerAddress" gorm:"not null;size:42;index"`
	NftAddress      string    `json:"nftAddress" gorm:"not null;size:42;index"`
	TokenId         string    `json:"tokenId" gorm:"not null;size:78;index"`
	DurationMinutes uint64    `json:"durationMinutes" gorm:"not null"`
	EndTime         time.Time `json:"endTime" gorm:"not null"`

	CreatedAt time.Time      `json:"createAt"`
	UpdatedAt time.Time      `json:"updateAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
