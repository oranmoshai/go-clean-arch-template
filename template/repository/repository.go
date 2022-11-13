package repository

import "github.com/oranmoshai/go-clean-arch-template/template/entity"

//TODO implement database methods

type Repository interface {
	ListUsers() (r entity.Result, err error)
}
