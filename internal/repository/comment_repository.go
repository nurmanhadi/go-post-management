package repository

import (
	"post-management/internal/entity"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}
func (r *CommentRepository) Save(comment *entity.Comment) error {
	return r.db.Save(comment).Error
}
func (r *CommentRepository) Delete(id int64) error {
	return r.db.Where("id = ?", id).Delete(&entity.Comment{}).Error
}
func (r *CommentRepository) CountById(id int64) (int64, error) {
	var count int64
	err := r.db.Model(&entity.Comment{}).Where("id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *CommentRepository) FindByPostIdAndUserId(postId, userId int64) (*entity.Comment, error) {
	comment := new(entity.Comment)
	err := r.db.Where("post_id = ? AND user_id = ?", postId, userId).First(comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}
