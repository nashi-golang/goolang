package models

type Person struct {
	Guid        string `json:"guid" gorm:"primaryKey"`
	WeddingGuid string `json:"wedding_guid" gorm:"type:varchar(36)"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}
