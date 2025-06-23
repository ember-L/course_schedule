package service

import (
	"course_schedule/internal/model"
	"course_schedule/internal/repository"
	"time"
)

type SystemConfigService interface {
	Get() (*SystemConfigWithWeek, error)
	Update(cfg *model.SystemConfig) error
}

type systemConfigService struct {
	repo repository.SystemConfigRepository
}

func NewSystemConfigService(repo repository.SystemConfigRepository) SystemConfigService {
	return &systemConfigService{repo: repo}
}

type SystemConfigWithWeek struct {
	model.SystemConfig
	CurrentWeek int `json:"currentWeek"`
}

func (s *systemConfigService) Get() (*SystemConfigWithWeek, error) {
	cfg, err := s.repo.Get()
	if err != nil {
		return nil, err
	}
	// 自动计算当前周次
	currentWeek := 1
	if cfg.StartDate != "" {
		start, err := time.Parse("2006-01-02", cfg.StartDate)
		if err == nil {
			days := int(time.Now().Sub(start).Hours() / 24)
			currentWeek = days/7 + 1
			if currentWeek < 1 {
				currentWeek = 1
			}
		}
	}
	return &SystemConfigWithWeek{
		SystemConfig: *cfg,
		CurrentWeek:  currentWeek,
	}, nil
}

func (s *systemConfigService) Update(cfg *model.SystemConfig) error {
	return s.repo.Update(cfg)
}
