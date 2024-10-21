package models

import "time"

type Wedding struct {
	Guid          string    `json:"guid" gorm:"primaryKey"`
	StartDatetime time.Time `json:"start_datetime"`
	Location      string    `json:"location"`
	Groom         string    `json:"groom"`
	Bride         string    `json:"bride"`
	Guests        []Person  `json:"guests" gorm:"foreignKey:WeddingGuid;references:Guid"`
}
