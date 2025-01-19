package entities

import "github.com/pborman/uuid"

type News struct {
	ID          string `gorm:"primary_key;type:varchar(36)" json:"id"`
	Link        string `gorm:"type:varchar(191);not null" json:"link"`
	Description string `gorm:"type:varchar(191);not null" json:"description"`
}

func NewNews(link, description string) *News {
	return &News{
		ID:          uuid.New(),
		Link:        link,
		Description: description,
	}
}