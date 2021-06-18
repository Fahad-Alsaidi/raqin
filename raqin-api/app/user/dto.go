package user

import "time"

type UserSignUp struct {
	FirstName string `json:"first_name" validate:"required,alpha"`
	LastName  string `json:"last_name" validate:"required,alpha"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,ascii"`
	Gender    string `json:"gender" validate:"required,alpha"`
}

type UserSignIn struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,ascii"`
}

type UserIDRequest struct {
	ID int `query:"id" validate:"required,number"`
}

type UpdateUserRequest struct {
	ID        int    `json:"id" validate:"required,number"`
	FirstName string `json:"first_name" validate:"required,alpha"`
	LastName  string `json:"last_name" validate:"required,alpha"`
	Email     string `json:"email" validate:"required,email"`
	Gender    string `json:"gender" validate:"required,alpha"`
	Role      string `json:"role" validate:"required,alpha"`
}

type UserResponse struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserByIDRequest struct {
	ID int `query:"id" validate:"required,number"`
}

type ChangePasswordRequest struct {
	ID          int    `query:"id" validate:"required,number"`
	OldPassword string `json:"old_password" validate:"required,ascii"`
	NewPassword string `json:"new_password" validate:"required,ascii"`
}
