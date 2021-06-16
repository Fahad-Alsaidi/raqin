package author

import (
	"raqin-api/storage/repo"
	"time"

	"raqin-api/utils/validator"
)

type AuthorService interface {
	NewAuthor(in NewAuthorRequest) (*AuthorResponse, error)
	DeleteAuthor(in ByID) error
	UpdateAuthor(in UpdateAuthorRequest) error
	AllAuthors() (*[]AuthorResponse, error)
	AuthorByID(in ByID) (*AuthorResponse, error)
}

type authorService struct {
	authorRepo AuthorRepo
}

func NewAuthorService(authorRepo AuthorRepo) *authorService {
	return &authorService{authorRepo}
}

func (auSrvc *authorService) NewAuthor(in NewAuthorRequest) (*AuthorResponse, error) {

	if err := validator.Validate(in); err != nil {
		return nil, err
	}

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

func (auSrvc *authorService) DeleteAuthor(in ByID) error {

	if err := validator.Validate(in); err != nil {
		return err
	}

	author := &repo.Author{
		ID:        in.ID,
		DeletedAt: time.Now(),
	}

	_, err := auSrvc.authorRepo.DeleteAuthor(author)
	if err != nil {
		return err
	}

	return nil
}

func (auSrvc *authorService) UpdateAuthor(in UpdateAuthorRequest) error {

	if err := validator.Validate(in); err != nil {
		return err
	}

	author := &repo.Author{
		ID:        in.ID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
	}

	_, err := auSrvc.authorRepo.UpdateAuthor(author)
	if err != nil {
		return err
	}

	return nil
}

func (auSrvc *authorService) AllAuthors() (*[]AuthorResponse, error) {

	authors, err := auSrvc.authorRepo.AllAuthors()
	if err != nil {
		return nil, err
	}

	authorsResponse := []AuthorResponse{}
	for _, author := range *authors {
		au := AuthorResponse{
			ID:        int64(author.ID),
			FirstName: author.FirstName,
			LastName:  author.LastName,
			CreatedAt: author.CreatedAt,
			UpdatedAt: author.UpdatedAt,
		}
		authorsResponse = append(authorsResponse, au)
	}

	return &authorsResponse, nil
}

func (auSrvc *authorService) AuthorByID(in ByID) (*AuthorResponse, error) {

	if err := validator.Validate(in); err != nil {
		return nil, err
	}

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
