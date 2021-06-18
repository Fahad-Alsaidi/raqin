package user

import (
	"context"
	"database/sql"
	"raqin-api/storage/repo"
	"raqin-api/utils/irror"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	errCanNotInsert    = irror.New("can not create user")
	errCanNotDelete    = irror.New("can not delete user")
	errCanNotUpdate    = irror.New("can not update user")
	errCanNotGetUser   = irror.New("can not get user")
	errCanNotGetUsers  = irror.New("can not get users")
	errCanNotPromote   = irror.New("can not promote user")
	errCanNotDemote    = irror.New("can not demote user")
	errCanNotChangePwd = irror.New("can not change password")
)

type UserRepo interface {
	SignUp(user *repo.User) (*repo.User, error)
	SignIn(user *repo.User) (*repo.User, error)
	DeleteUser(user *repo.User) (int64, error)
	UpdateUser(user *repo.User) (int64, error)
	AllUser() (repo.UserSlice, error)
	UserByID(id int) (*repo.User, error)
	PromoteUser(id int) (int64, error)
	DemoteUser(id int) (int64, error)
	ChangePassword(user *repo.User) (int64, error)
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
		return nil, errCanNotInsert.Wrap(err)
	}

	err = user.Insert(ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return nil, errCanNotInsert.Wrap(err)
	}

	tx.Commit()
	return user, nil
}

func (usrepo *userRepo) SignIn(user *repo.User) (*repo.User, error) {

	ctx := context.Background()
	var err error

	user, err = repo.Users(qm.Where("email=?", user.Email)).One(ctx, usrepo.db)
	if err != nil {
		return nil, errCanNotGetUser.Wrap(err)
	}

	return user, nil
}

func (usrepo *userRepo) DeleteUser(user *repo.User) (int64, error) {

	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotDelete.Wrap(err)
	}

	n, err := user.Delete(ctx, tx)
	if err != nil {
		tx.Rollback()
		return 0, errCanNotDelete.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (usrepo *userRepo) UpdateUser(user *repo.User) (int64, error) {

	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotUpdate.Wrap(err)
	}

	n, err := user.Update(ctx, tx, boil.Blacklist("id", "password", "role", "created_at", "deleted_at"))
	if err != nil {
		tx.Rollback()
		return 0, errCanNotUpdate.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (usrepo *userRepo) AllUser() (repo.UserSlice, error) {

	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetUsers.Wrap(err)
	}

	users, err := repo.Users().All(ctx, tx)
	if err != nil {
		return nil, errCanNotGetUsers.Wrap(err)
	}

	return users, nil
}

func (usrepo *userRepo) UserByID(id int) (*repo.User, error) {

	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errCanNotGetUser.Wrap(err)
	}

	user, err := repo.FindUser(ctx, tx, id)
	if err != nil {
		return nil, errCanNotGetUser.Wrap(err)
	}

	return user, nil
}

func (usrepo *userRepo) PromoteUser(id int) (int64, error) {
	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotPromote.Wrap(err)
	}

	user, err := repo.FindUser(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return 0, errCanNotPromote.Wrap(err)
	}

	user.Role = repo.UserRoleRAQIN

	n, err := user.Update(ctx, tx, boil.Whitelist("role"))
	if err != nil {
		tx.Rollback()
		return 0, errCanNotPromote.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (usrepo *userRepo) ChangePassword(user *repo.User) (int64, error) {
	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotChangePwd.Wrap(err)
	}

	n, err := user.Update(ctx, tx, boil.Whitelist("password"))
	if err != nil {
		tx.Rollback()
		return 0, errCanNotChangePwd.Wrap(err)
	}

	tx.Commit()
	return n, nil
}

func (usrepo *userRepo) DemoteUser(id int) (int64, error) {
	ctx := context.Background()
	tx, err := usrepo.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, errCanNotDemote.Wrap(err)
	}

	user, err := repo.FindUser(ctx, tx, id)
	if err != nil {
		tx.Rollback()
		return 0, errCanNotDemote.Wrap(err)
	}

	user.Role = repo.UserRoleVOLUNTEER

	n, err := user.Update(ctx, tx, boil.Whitelist("role"))
	if err != nil {
		tx.Rollback()
		return 0, errCanNotDemote.Wrap(err)
	}

	tx.Commit()
	return n, nil
}
