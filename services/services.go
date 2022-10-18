package services

import (
	"book-api-gateway/config"
	"book-api-gateway/genproto/book_service"
	"fmt"

	"google.golang.org/grpc"
)

type ServicesI interface {
	BookCategoryService() book_service.BookCategoryServiceClient
	BookService() book_service.BookServiceClient
}

type servicesRepo struct {
	bookCategoryService book_service.BookCategoryServiceClient
	bookService         book_service.BookServiceClient
}

func NewServicesRepo(c *config.Config) (ServicesI, error) {
	connBookCategoryService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", c.BookServiceHost, c.BookServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &servicesRepo{
		bookCategoryService: book_service.NewBookCategoryServiceClient(connBookCategoryService),
		bookService:         book_service.NewBookServiceClient(connBookCategoryService),
	}, nil

}

func (s *servicesRepo) BookCategoryService() book_service.BookCategoryServiceClient {
	return s.bookCategoryService
}
func (s *servicesRepo) BookService() book_service.BookServiceClient {
	return s.bookService
}
