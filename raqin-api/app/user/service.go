package user

import (
	"raqin-api/storage/repo"
	"time"

	"github.com/volatiletech/null/v8"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(in UserSignUp) (*UserResponse, error)
	SignIn(in UserSignIn) (*UserResponse, error)
	DeleteUser(in UserIDRequest) error
	UpdateUser(in UpdateUserRequest) (*UserResponse, error)
	AllUsers() ([]UserResponse, error)
	UserByID(in GetUserByIDRequest) (*UserResponse, error)
	PromoteUser(in UserIDRequest) error
	DemoteUser(in UserIDRequest) error
	ChangePassword(in ChangePasswordRequest) error
}

type userService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *userService {
	return &userService{userRepo}
}

func (usSrvc *userService) SignUp(in UserSignUp) (*UserResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), 8)
	if err != nil {
		return nil, err
	}

	user := &repo.User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  string(hashedPassword),
		Role:      repo.UserRoleVOLUNTEER,
		Gender:    null.StringFrom(in.Gender),
	}

	c, err := usSrvc.userRepo.SignUp(user)
	if err != nil {
		return nil, err
	}

	res := &UserResponse{
		ID: int64(c.ID),
	}
	return res, nil

}

func (usSrvc *userService) SignIn(in UserSignIn) (*UserResponse, error) {

	user := &repo.User{
		Email: in.Email,
	}

	usr, err := usSrvc.userRepo.SignIn(user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(in.Password))
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
		Gender:    usr.Gender.String,
		Role:      usr.Role,
	}, nil

}

func (usSrvc *userService) DeleteUser(in UserIDRequest) error {

	user := &repo.User{
		ID: int(in.ID),
	}

	n, err := usSrvc.userRepo.DeleteUser(user)
	if err != nil || n == 0 {
		return err
	}

	return nil

}

func (usSrvc *userService) UpdateUser(in UpdateUserRequest) (*UserResponse, error) {

	user := &repo.User{
		ID:        in.ID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Gender:    null.StringFrom(in.Gender),
		UpdatedAt: time.Now(),
	}

	ca, err := usSrvc.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:        int64(ca.ID),
		FirstName: ca.FirstName,
		LastName:  ca.LastName,
		Email:     ca.Email,
		Gender:    ca.Gender.String,
		Role:      ca.Role,
		CreatedAt: ca.CreatedAt,
		UpdatedAt: ca.UpdatedAt,
	}, nil

}

func (usSrvc *userService) AllUsers() ([]UserResponse, error) {

	users, err := usSrvc.userRepo.AllUser()
	if err != nil {
		return nil, err
	}

	usersResponse := []UserResponse{}
	for _, user := range users {
		us := UserResponse{
			ID:        int64(user.ID),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Gender:    user.Gender.String,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		usersResponse = append(usersResponse, us)
	}

	return usersResponse, nil

}

func (usSrvc *userService) UserByID(in GetUserByIDRequest) (*UserResponse, error) {

	user, err := usSrvc.userRepo.UserByID(in.ID)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:        int64(user.ID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Gender:    user.Gender.String,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil

}

func (usSrvc *userService) PromoteUser(in UserIDRequest) error {

	_, err := usSrvc.userRepo.PromoteUser(in.ID)
	if err != nil {
		return err
	}

	return nil
}

func (usSrvc *userService) ChangePassword(in ChangePasswordRequest) error {

	user, err := usSrvc.userRepo.UserByID(in.ID)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.OldPassword))
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), 8)
	if err != nil {
		return err
	}

	newUser := &repo.User{
		ID:       in.ID,
		Password: string(hashedPassword),
	}

	_, err = usSrvc.userRepo.ChangePassword(newUser)
	if err != nil {
		return err
	}

	return nil
}

func (usSrvc *userService) DemoteUser(in UserIDRequest) error {

	_, err := usSrvc.userRepo.DemoteUser(in.ID)
	if err != nil {
		return err
	}

	return nil
}
