package models

type People struct {
	WeddingGuid string `json:"wedding-guid" gorm:"primaryKey"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone-number"`
}
