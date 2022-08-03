package store

import "kun-blog-golang/pkg/store/models"

type StoreInterface interface {
	List() []*models.Post
	DeleteBySlug(slug string) error
	GetBySlug(slug string) (*models.Post, error)
	Save(post *models.Post) error
}
