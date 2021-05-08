package user

import (
	"context"
	"database/sql"
	"errors"
	"raqin-api/storage/repo"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type UserRepo interface {
	SignUp(user *repo.User) (*repo.User, error)
	SignIn(user *repo.User) (*repo.User, error)
	DeleteUser(user *repo.User) (int64, error)
	UpdateUser(user *repo.User) (*repo.User, error)
	AllUser() (repo.UserSlice, error)
	UserByID(id int) (*repo.User, error)
	PromoteUser(id int) (*repo.User, error)
	DemoteUser(id int) (*repo.User, error)
	ChangePassword(user *repo.User) (*repo.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{db}
}

func (usrepo *userRepo) SignUp(user *repo.User) (*repo.User, error) {

	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	err = user.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return user, nil

}

func (usrepo *userRepo) SignIn(user *repo.User) (*repo.User, error) {

	ctx := context.Background()
	var err error

	user, err = repo.Users(qm.Where("email=?", user.Email)).One(ctx, usrepo.db)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (usrepo *userRepo) DeleteUser(user *repo.User) (int64, error) {

	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	n, err := user.Delete(ctx, tx)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	return n, nil

}

func (usrepo *userRepo) UpdateUser(user *repo.User) (*repo.User, error) {

	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	n, err := user.Update(ctx, tx, boil.Blacklist("id", "password", "role", "created_at", "deleted_at"))
	if err != nil || n == 0 {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return user, nil

}

func (usrepo *userRepo) AllUser() (repo.UserSlice, error) {

	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	users, err := repo.Users().All(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return users, nil

}

func (usrepo *userRepo) UserByID(id int) (*repo.User, error) {

	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	user, err := repo.FindUser(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return user, nil

}

func (usrepo *userRepo) PromoteUser(id int) (*repo.User, error) {
	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	user, err := repo.FindUser(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	user.Role = repo.UserRoleRAQIN

	n, err := user.Update(ctx, tx, boil.Whitelist("role"))
	if err != nil || n == 0 {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return user, nil
}

func (usrepo *userRepo) ChangePassword(user *repo.User) (*repo.User, error) {
	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	n, err := user.Update(ctx, tx, boil.Whitelist("password"))
	if err != nil || n == 0 {
		tx.Rollback()
		return nil, errors.New("could not update password")
	}

	tx.Commit()

	return user, nil
}

func (usrepo *userRepo) DemoteUser(id int) (*repo.User, error) {
	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	user, err := repo.FindUser(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	user.Role = repo.UserRoleVOLUNTEER

	n, err := user.Update(ctx, tx, boil.Whitelist("role"))
	if err != nil || n == 0 {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return user, nil
}
