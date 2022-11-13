package service

import (
	"github.com/oranmoshai/go-clean-arch-template/template/repository"
	"github.com/oranmoshai/go-clean-arch-template/template/usecase"
)

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Run() (err error) {
	u := usecase.NewSampleUseCase(s.repository)
	err = u.SampleUseCase()
	if err != nil {
		return err
	}

	return nil
}
