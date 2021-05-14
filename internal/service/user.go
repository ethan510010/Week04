package service

import (
	v1 "github.com/ethan510010/Week_04/api/book/v1"
	"context"
	"github.com/ethan510010/Week_04/internal/biz"
)


type UserService struct {
	b *biz.BookUsecase
}

func NewUserService(u *biz.BookUsecase) v1.BookServiceServer {
	return &UserService{b: u}
}

func (u *UserService) GetBook(ctx context.Context, request *v1.BookRequest) (*v1.BookInfo, error) {
	book := u.b.GetBookInfo(request.BookId)
	return &v1.BookInfo{
		BookId: book.BookId,
		BookName: book.BookName,
	}, nil
}