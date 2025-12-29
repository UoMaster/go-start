package dto

import "errors"

const (
	MESSAGE_FAILED_GET_DATA_FROM_BODY = "failed get data from body"
	MESSAGE_FAILED_REGISTER_USER      = "failed create user"
	MESSAGE_FAILED_GET_USER           = "failed get user"
	MESSAGE_FAILED_LOGIN              = "failed login"
	MESSAGE_FAILED_UPDATE_USER        = "failed update user"
	MESSAGE_FAILED_DELETE_USER        = "failed delete user"
	MESSAGE_FAILED_TOKEN_NOT_VALID    = "token not valid"
	MESSAGE_FAILED_TOKEN_NOT_FOUND    = "token not found"
	MESSAGE_FAILED_PROSES_REQUEST     = "failed proses request"
	MESSAGE_FAILED_DENIED_ACCESS      = "denied access"

	MESSAGE_SUCCESS_REGISTER_USER = "success create user"
	MESSAGE_SUCCESS_GET_USER      = "success get user"
	MESSAGE_SUCCESS_LOGIN         = "success login"
	MESSAGE_SUCCESS_UPDATE_USER   = "success update user"
	MESSAGE_SUCCESS_DELETE_USER   = "success delete user"
)

var (
	ErrCreateUser         = errors.New("failed to create user")
	ErrGetUserById        = errors.New("failed to get user by id")
	ErrGetUserByEmail     = errors.New("failed to get user by email")
	ErrEmailAlreadyExists = errors.New("email already exist")
	ErrUpdateUser         = errors.New("failed to update user")
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailNotFound      = errors.New("email not found")
	ErrDeleteUser         = errors.New("failed to delete user")
	ErrTokenInvalid       = errors.New("token invalid")
	ErrTokenExpired       = errors.New("token expired")
)

type (
	UserCreateRequest struct {
		Name       string `json:"name" form:"name" binding:"required,min=2,max=100"`
		TelpNumber string `json:"telp_number" form:"telp_number" binding:"omitempty,min=8,max=20"`
		Email      string `json:"email" form:"email" binding:"required,email"`
		Password   string `json:"password" form:"password" binding:"required,min=8"`
	}

	UserResponse struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		TelpNumber string `json:"telp_number"`
		Role       string `json:"role"`
		ImageUrl   string `json:"image_url"`
		IsVerified bool   `json:"is_verified"`
	}

	UserUpdateRequest struct {
		Name       string `json:"name" form:"name" binding:"omitempty,min=2,max=100"`
		TelpNumber string `json:"telp_number" form:"telp_number" binding:"omitempty,min=8,max=20"`
		Email      string `json:"email" form:"email" binding:"omitempty,email"`
	}

	UserUpdateResponse struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		TelpNumber string `json:"telp_number"`
		Role       string `json:"role"`
		Email      string `json:"email"`
		IsVerified bool   `json:"is_verified"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
)
