package biz

type Book struct {
	BookId int32
	BookName string
}

type BookRepo interface {
	GetBookInfo(id int32) Book
}

type BookUsecase struct {
	repo BookRepo
}

func NewBookUsecase(repo BookRepo) *BookUsecase {
	return &BookUsecase{
		repo: repo,
	}
}

func (b *BookUsecase) GetBookInfo(id int32) Book {
	return b.repo.GetBookInfo(id)
}
