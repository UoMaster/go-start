package dto

import "errors"

const (
	MESSAGE_FAILED_REFRESH_TOKEN  = "failed to refresh token"
	MESSAGE_SUCCESS_REFRESH_TOKEN = "success refresh token"
)

var (
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrRefreshTokenNotFound = errors.New("refresh token not found")
)

type (
	TokenResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		Role         string `json:"role"`
	}

	RefreshTokenRequest struct {
		RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required"`
	}
)
