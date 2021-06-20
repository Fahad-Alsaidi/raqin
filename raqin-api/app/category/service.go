package category

import (
	"raqin-api/storage/repo"
	"raqin-api/utils/validator"
	"time"
)

type CategoryService interface {
	NewCategory(in NewCategoryRequest) (*CategoryResponse, error)
	DeleteCategory(in DeleteCategoryRequest) error
	UpdateCategory(in UpdateCategoryRequest) error
	AllCategories() ([]CategoryResponse, error)
}

type categoryService struct {
	categoryRepo CategoryRepo
}

func NewCategoryService(categoryRepo CategoryRepo) *categoryService {
	return &categoryService{categoryRepo}
}

func (catSrvc *categoryService) NewCategory(in NewCategoryRequest) (*CategoryResponse, error) {

	if err := validator.Validate(in); err != nil {
		return nil, err
	}

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

	if err := validator.Validate(in); err != nil {
		return err
	}

	category := &repo.Category{
		ID: int(in.ID),
	}

	_, err := catSrvc.categoryRepo.DeleteCategory(category)
	if err != nil {
		return err
	}

	return nil
}

func (catSrvc *categoryService) UpdateCategory(in UpdateCategoryRequest) error {

	if err := validator.Validate(in); err != nil {
		return err
	}

	category := &repo.Category{
		ID:        in.ID,
		Name:      in.Name,
		UpdatedAt: time.Now(),
	}

	_, err := catSrvc.categoryRepo.UpdateCategory(category)
	if err != nil {
		return err
	}

	return nil
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
