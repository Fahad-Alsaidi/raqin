package category

import (
	"context"
	"database/sql"
	"raqin-api/storage/repo"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CategoryRepo interface {
	NewCategory(category *repo.Category) (*repo.Category, error)
	DeleteCategory(category *repo.Category) (int64, error)
	UpdateCategory(category *repo.Category) (*repo.Category, error)
	AllCategory() (repo.CategorySlice, error)
	CategoryByID(id int) (*repo.Category, error)
}

type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) *categoryRepo {
	return &categoryRepo{db}
}

func (carepo *categoryRepo) NewCategory(category *repo.Category) (*repo.Category, error) {

	ctx := context.Background()
	tx, err := carepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = category.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return category, nil

}

func (carepo *categoryRepo) DeleteCategory(category *repo.Category) (int64, error) {

	ctx := context.Background()
	tx, err := carepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	n, err := category.Delete(ctx, tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return n, nil

}

func (carepo *categoryRepo) UpdateCategory(category *repo.Category) (*repo.Category, error) {

	ctx := context.Background()
	tx, err := carepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	n, err := category.Update(ctx, tx, boil.Infer())
	if err != nil || n == 0 {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return category, nil

}

func (carepo *categoryRepo) AllCategory() (repo.CategorySlice, error) {

	ctx := context.Background()
	tx, err := carepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	categories, err := repo.Categories().All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return categories, nil

}

func (carepo *categoryRepo) CategoryByID(id int) (*repo.Category, error) {

	ctx := context.Background()
	tx, err := carepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	category, err := repo.FindCategory(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return category, nil

}
