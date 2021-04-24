package author

import (
	"raqin-api/storage/repo"
)

type AuthorService interface {
	NewAuthor(in NewAuthorRequest) (NewAuthorResponse, error)
}

type authorService struct {
	authorRepo AuthorRepo
}

func NewBookService(authorRepo AuthorRepo) *authorService {
	return &authorService{authorRepo}
}

func (auSrvc *authorService) NewAuthor(in NewAuthorRequest) (res NewAuthorResponse, err error) {

	author := &repo.Author{
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}

	a, err := auSrvc.authorRepo.NewAuthor(author)
	if err != nil {
		return res, err
	}

	res = NewAuthorResponse{
		FirstName: a.FirstName,
		LastName:  a.LastName,
	}
	return res, nil

}
