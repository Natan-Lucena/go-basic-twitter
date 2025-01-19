package services

import (
	"crud-go/internal/entities"
	"crud-go/internal/repositories"
	errors "crud-go/pkg/err"
)

type NewsService struct {
	newsRepository *repositories.NewsRepository
}

func NewNewsService() *NewsService {
	newsRepository := repositories.NewNewsRepository()
	return &NewsService{
		newsRepository: newsRepository,
	}
}

func (service *NewsService) CreateNews(description, link string) error {
	news := entities.NewNews(link, description)
	_, err := service.newsRepository.Save(news)
	if err != nil {
		return errors.ErrWhileSavingNews
	}
	return nil
}