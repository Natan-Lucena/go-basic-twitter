package repositories

import (
	"crud-go/internal/entities"
	GORM "crud-go/pkg/gorm"

	"gorm.io/gorm"
)

type NewsRepository struct {
	db *gorm.DB
}

type NewsFilters struct {
	ID          *string
	Link        *string
	Description *string
}

type Pagination struct {
	Page  int 
	Limit int 
}

func NewNewsRepository() *NewsRepository{
	db, _ := GORM.InitDB()
	return &NewsRepository{
		db: db,
	}
}

func (repository *NewsRepository) Save(news *entities.News) (*entities.News, error) {
	if err := repository.db.Save(news).Error; err != nil {
		return nil, err
	}
	return news, nil
}

func (repository *NewsRepository) List(filters *NewsFilters, pagination *Pagination) ([]entities.News, error) {
	var newsList []entities.News

	query := repository.db.Model(&entities.News{})

	if filters != nil {
		if filters.ID != nil {
			query = query.Where("id = ?", *filters.ID)
		}
		if filters.Link != nil {
			query = query.Where("link LIKE ?", "%"+*filters.Link+"%")
		}
		if filters.Description != nil {
			query = query.Where("description LIKE ?", "%"+*filters.Description+"%")
		}
	}

	if pagination != nil {
		offset := (pagination.Page - 1) * pagination.Limit
		query = query.Offset(offset).Limit(pagination.Limit)
	}

	if err := query.Find(&newsList).Error; err != nil {
		return nil, err
	}

	return newsList, nil
}