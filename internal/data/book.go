package data

import "github.com/ethan510010/Week_04/internal/biz"

type bookRepo struct {}

func (b *bookRepo) GetBookInfo(id int32) biz.Book  {
	return biz.Book{
		BookId: 1,
		BookName: "Go Concurrency",
	}
}

func NewBookRepo() biz.BookRepo {
	return &bookRepo{}
}