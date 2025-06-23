package repository

import (
	"course_schedule/internal/model"
)

// AnnouncementRepository 公告仓库接口
// 支持CRUD
type AnnouncementRepository interface {
	GetAll() ([]model.Announcement, error)
	GetByID(id uint) (*model.Announcement, error)
	Create(announcement *model.Announcement) error
	Update(announcement *model.Announcement) error
	Delete(id uint) error
}

type announcementRepository struct {
	db *Database
}

func NewAnnouncementRepository(db *Database) AnnouncementRepository {
	return &announcementRepository{db: db}
}

func (r *announcementRepository) GetAll() ([]model.Announcement, error) {
	var list []model.Announcement
	result := r.db.DB.Order("created_at desc").Find(&list)
	return list, result.Error
}

func (r *announcementRepository) GetByID(id uint) (*model.Announcement, error) {
	var a model.Announcement
	result := r.db.DB.First(&a, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &a, nil
}

func (r *announcementRepository) Create(a *model.Announcement) error {
	return r.db.DB.Create(a).Error
}

func (r *announcementRepository) Update(a *model.Announcement) error {
	return r.db.DB.Save(a).Error
}

func (r *announcementRepository) Delete(id uint) error {
	return r.db.DB.Delete(&model.Announcement{}, id).Error
}
