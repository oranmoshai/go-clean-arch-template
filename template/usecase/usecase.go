package usecase

import "github.com/oranmoshai/go-clean-arch-template/template/repository"

type SampleUseCase struct {
	repository repository.Repository
}

func NewSampleUseCase(repo repository.Repository) *SampleUseCase {
	return &SampleUseCase{
		repository: repo,
	}
}

func (s *SampleUseCase) SampleUseCase() error {

	_, err := s.repository.ListUsers()
	if err != nil {
		return err
	}

	return nil
}
