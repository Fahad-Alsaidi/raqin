package category

import (
	"context"
	"database/sql"
	"raqin-api/storage/repo"
	"raqin-api/utils/irror"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CategoryRepo interface {
	NewCategory(category *repo.Category) (*repo.Category, error)
	DeleteCategory(category *repo.Category) (int64, error)
	UpdateCategory(category *repo.Category) (int64, error)
	AllCategory() (repo.CategorySlice, error)
	CategoryByID(id int) (*repo.Category, error)
}

var (
	errCanNotInsert        = irror.New("can not insert category")
	errCanNotDelete        = irror.New("can not delete category")
	errCanNotUpdate        = irror.New("can not update category")
	errCanNotGetCategory   = irror.New("can not get category")
	errCanNotGetCategories = irror.New("can not get categories")
)

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
		return nil, errCanNotInsert.Wrap(err)
	}

	err = category.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, errCanNotInsert.Wrap(err)
	}

	tx.Commit()
	return category, nil
}

func (carepo *categoryRepo) DeleteCategory(category *repo.Category) (int64, error) {

	ctx := context.Background()
	tx, err := carepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotDelete.Wrap(err)
	}

	n, err := category.Delete(ctx, tx)
	if err != nil {
		tx.Rollback()
		return 0, errCanNotDelete.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (carepo *categoryRepo) UpdateCategory(category *repo.Category) (int64, error) {

	ctx := context.Background()
	tx, err := carepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotUpdate.Wrap(err)
	}

	n, err := category.Update(ctx, tx, boil.Whitelist("name", "updated_at"))
	if err != nil {
		tx.Rollback()
		return 0, errCanNotUpdate.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (carepo *categoryRepo) AllCategory() (repo.CategorySlice, error) {

	ctx := context.Background()
	tx, err := carepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetCategories.Wrap(err)
	}

	categories, err := repo.Categories(
		qm.Where("deleted_at = '0000-00-00 00:00:00'")).
		All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, errCanNotGetCategories.Wrap(err)
	}

	return categories, nil
}

func (carepo *categoryRepo) CategoryByID(id int) (*repo.Category, error) {

	ctx := context.Background()
	tx, err := carepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetCategory.Wrap(err)
	}

	category, err := repo.FindCategory(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return nil, errCanNotGetCategory.Wrap(err)
	}

	return category, nil
}
