package repository

import (
	"post-management/internal/entity"

	"gorm.io/gorm"
)

type LikeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{
		db: db,
	}
}
func (r *LikeRepository) Save(like *entity.Like) error {
	return r.db.Save(like).Error
}
func (r *LikeRepository) Delete(id int64) error {
	return r.db.Where("id = ?", id).Delete(&entity.Like{}).Error
}
func (r *LikeRepository) CountById(id int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Like{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *LikeRepository) CountByPostIdAndUserId(postId, userId int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Like{}).Where("post_id = ? AND user_id = ?", postId, userId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *LikeRepository) FindByPostIdAndUserId(postId, userId int64) (*entity.Like, error) {
	like := new(entity.Like)
	err := r.db.Where("post_id = ? AND user_id = ?", postId, userId).First(like).Error
	if err != nil {
		return nil, err
	}
	return like, nil
}
