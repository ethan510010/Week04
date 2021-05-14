// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/ethan510010/Week_04/internal/biz"
	"github.com/ethan510010/Week_04/internal/data"
)

// Injectors from wire.go:

func InitBookUsecase() *biz.BookUsecase {
	bookRepo := data.NewBookRepo()
	bookUsecase := biz.NewBookUsecase(bookRepo)
	return bookUsecase
}
