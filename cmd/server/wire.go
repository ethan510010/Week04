//+build wireinject

package main

import (
	"github.com/ethan510010/Week_04/internal/biz"
	"github.com/ethan510010/Week_04/internal/data"

	"github.com/google/wire"
)

func InitBookUsecase() *biz.BookUsecase {
	wire.Build(biz.NewBookUsecase, data.NewBookRepo)
	return &biz.BookUsecase{}
}


