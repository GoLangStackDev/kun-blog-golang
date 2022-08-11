package store

import "kun-blog-golang/pkg/store/models"

type StoreInterface interface {
	List() (posts []*models.Post, err error) //调整
	DeleteBySlug(slug string) error
	GetBySlug(slug string) (*models.Post, error)
	Save(post *models.Post) error
}
