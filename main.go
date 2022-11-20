package main

import (
	"fmt"
	"github.com/oranmoshai/go-clean-arch-template/template/repository/mockdatabase"
	"github.com/oranmoshai/go-clean-arch-template/template/service"
)

func main() {
	fmt.Println("Started")
	repo := mockdatabase.New()
	s := service.New(repo)
	err := s.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success...")
}
