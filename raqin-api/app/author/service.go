package author

import (
	"raqin-api/storage/repo"
)

type AuthorService interface {
	NewAuthor(in NewAuthorRequest) (*AuthorResponse, error)
	DeleteAuthor(in DeleteAuthorRequest) error
	UpdateAuthor(in UpdateAuthorRequest) (*AuthorResponse, error)
	AllAuthors() ([]AuthorResponse, error)
	AuthorByID(in GetAuthorByIDRequest) (*AuthorResponse, error)
}

type authorService struct {
	authorRepo AuthorRepo
}

func NewAuthorService(authorRepo AuthorRepo) *authorService {
	return &authorService{authorRepo}
}

func (auSrvc *authorService) NewAuthor(in NewAuthorRequest) (*AuthorResponse, error) {

	author := &repo.Author{
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}

	a, err := auSrvc.authorRepo.NewAuthor(author)
	if err != nil {
		return nil, err
	}

	res := &AuthorResponse{
		ID:        int64(a.ID),
		FirstName: a.FirstName,
		LastName:  a.LastName,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
	return res, nil

}

func (auSrvc *authorService) DeleteAuthor(in DeleteAuthorRequest) error {

	author := &repo.Author{
		ID: in.ID,
	}

	n, err := auSrvc.authorRepo.DeleteAuthor(author)
	if err != nil || n == 0 {
		return err
	}

	return nil

}

func (auSrvc *authorService) UpdateAuthor(in UpdateAuthorRequest) (*AuthorResponse, error) {

	author := &repo.Author{
		ID:        in.ID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}

	au, err := auSrvc.authorRepo.UpdateAuthor(author)
	if err != nil {
		return nil, err
	}

	return &AuthorResponse{
		ID:        int64(au.ID),
		FirstName: au.FirstName,
		LastName:  au.LastName,
		CreatedAt: au.CreatedAt,
		UpdatedAt: au.UpdatedAt,
	}, nil

}

func (auSrvc *authorService) AllAuthors() ([]AuthorResponse, error) {

	authors, err := auSrvc.authorRepo.AllAuthors()
	if err != nil {
		return nil, err
	}

	authorsResponse := []AuthorResponse{}
	for _, author := range authors {
		au := AuthorResponse{
			ID:        int64(author.ID),
			FirstName: author.FirstName,
			LastName:  author.LastName,
			CreatedAt: author.CreatedAt,
			UpdatedAt: author.UpdatedAt,
		}
		authorsResponse = append(authorsResponse, au)
	}

	return authorsResponse, nil

}

func (auSrvc *authorService) AuthorByID(in GetAuthorByIDRequest) (*AuthorResponse, error) {

	author, err := auSrvc.authorRepo.AuthorByID(in.ID)
	if err != nil {
		return nil, err
	}

	return &AuthorResponse{
		ID:        int64(author.ID),
		FirstName: author.FirstName,
		LastName:  author.LastName,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}, nil

}
