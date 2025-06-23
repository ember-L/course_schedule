package service

import (
	"course_schedule/internal/model"
	"course_schedule/internal/repository"
)

type AnnouncementService interface {
	GetAll() ([]model.Announcement, error)
	GetByID(id uint) (*model.Announcement, error)
	Create(a *model.Announcement) error
	Update(a *model.Announcement) error
	Delete(id uint) error
}

type announcementService struct {
	repo repository.AnnouncementRepository
}

func NewAnnouncementService(repo repository.AnnouncementRepository) AnnouncementService {
	return &announcementService{repo: repo}
}

func (s *announcementService) GetAll() ([]model.Announcement, error) {
	return s.repo.GetAll()
}
func (s *announcementService) GetByID(id uint) (*model.Announcement, error) {
	return s.repo.GetByID(id)
}
func (s *announcementService) Create(a *model.Announcement) error {
	return s.repo.Create(a)
}
func (s *announcementService) Update(a *model.Announcement) error {
	return s.repo.Update(a)
}
func (s *announcementService) Delete(id uint) error {
	return s.repo.Delete(id)
}
