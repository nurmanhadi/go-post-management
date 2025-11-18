package repository

import (
	"post-management/internal/entity"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}
func (r *PostRepository) Save(post *entity.Post) error {
	return r.db.Save(post).Error
}
func (r *PostRepository) FindById(id int64) (*entity.Post, error) {
	post := new(entity.Post)
	err := r.db.Where("id = ?", id).First(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (r *PostRepository) Delete(id int64) error {
	return r.db.Where("id = ?", id).Delete(&entity.Post{}).Error
}
func (r *PostRepository) CountById(id int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Post{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
