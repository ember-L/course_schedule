package repository

import (
	"course_schedule/internal/model"
)

type SystemConfigRepository interface {
	Get() (*model.SystemConfig, error)
	Update(cfg *model.SystemConfig) error
}

type systemConfigRepository struct {
	db *Database
}

func NewSystemConfigRepository(db *Database) SystemConfigRepository {
	return &systemConfigRepository{db: db}
}

func (r *systemConfigRepository) Get() (*model.SystemConfig, error) {
	var cfg model.SystemConfig
	result := r.db.DB.First(&cfg, 1)
	return &cfg, result.Error
}

func (r *systemConfigRepository) Update(cfg *model.SystemConfig) error {
	cfg.ID = 1
	return r.db.DB.Save(cfg).Error
}
