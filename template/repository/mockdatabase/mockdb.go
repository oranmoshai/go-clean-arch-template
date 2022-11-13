package mockdatabase

import "github.com/oranmoshai/go-clean-arch-template/template/entity"

type MockDB struct {
}

func NewMockDB() *MockDB {
	return &MockDB{}
}

func (m *MockDB) ListUsers() (r entity.Result, err error) {
	return r, nil
}
