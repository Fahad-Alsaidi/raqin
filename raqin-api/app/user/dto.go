package user

import "time"

type UserSignUp struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserIDRequest struct {
	ID int `query:"id"`
}

type UpdateUserRequest struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Role      string `json:"role"`
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
	ID int `query:"id"`
}

type ChangePasswordRequest struct {
	ID          int    `query:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
