package category

import (
	"raqin-api/storage/repo"
)

type CategoryService interface {
	NewCategory(in NewCategoryRequest) (*CategoryResponse, error)
	DeleteCategory(in DeleteCategoryRequest) error
	UpdateCategory(in UpdateCategoryRequest) (*CategoryResponse, error)
	AllCategories() ([]CategoryResponse, error)
	CategoryByID(in GetCategoryByIDRequest) (*CategoryResponse, error)
}

type categoryService struct {
	categoryRepo CategoryRepo
}

func NewCategoryService(categoryRepo CategoryRepo) *categoryService {
	return &categoryService{categoryRepo}
}

func (catSrvc *categoryService) NewCategory(in NewCategoryRequest) (*CategoryResponse, error) {

	category := &repo.Category{
		Name: in.Name,
	}

	c, err := catSrvc.categoryRepo.NewCategory(category)
	if err != nil {
		return nil, err
	}

	res := &CategoryResponse{
		ID:        int64(c.ID),
		Name:      c.Name,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
	return res, nil

}

func (catSrvc *categoryService) DeleteCategory(in DeleteCategoryRequest) error {

	category := &repo.Category{
		ID: int(in.ID),
	}

	n, err := catSrvc.categoryRepo.DeleteCategory(category)
	if err != nil || n == 0 {
		return err
	}

	return nil

}

func (catSrvc *categoryService) UpdateCategory(in UpdateCategoryRequest) (*CategoryResponse, error) {

	category := &repo.Category{
		ID:   in.ID,
		Name: in.Name,
	}

	ca, err := catSrvc.categoryRepo.UpdateCategory(category)
	if err != nil {
		return nil, err
	}

	return &CategoryResponse{
		ID:        int64(ca.ID),
		Name:      ca.Name,
		CreatedAt: ca.CreatedAt,
		UpdatedAt: ca.UpdatedAt,
	}, nil

}

func (catSrvc *categoryService) AllCategories() ([]CategoryResponse, error) {

	categories, err := catSrvc.categoryRepo.AllCategory()
	if err != nil {
		return nil, err
	}

	categoriesResponse := []CategoryResponse{}
	for _, category := range categories {
		ca := CategoryResponse{
			ID:        int64(category.ID),
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}
		categoriesResponse = append(categoriesResponse, ca)
	}

	return categoriesResponse, nil

}

func (catSrvc *categoryService) CategoryByID(in GetCategoryByIDRequest) (*CategoryResponse, error) {

	category, err := catSrvc.categoryRepo.CategoryByID(in.ID)
	if err != nil {
		return nil, err
	}

	return &CategoryResponse{
		ID:        int64(category.ID),
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil

}
