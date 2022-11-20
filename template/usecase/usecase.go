package usecase

import "github.com/oranmoshai/go-clean-arch-template/template/repository"

type Usecase struct {
	repository repository.Repository
}

func New(repo repository.Repository) *Usecase {
	return &Usecase{
		repository: repo,
	}
}

func (u *Usecase) Sample() error {

	_, err := u.repository.ListUsers()
	if err != nil {
		return err
	}

	return nil
}
