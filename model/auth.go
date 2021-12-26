package model

import "github.com/golang-jwt/jwt"

type GetTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AppKey   string `json:"app_key"`
}

type GetTokenResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
	Data    Token  `json:"data,omitempty"`
}

type Token struct {
	Token string `json:"token"`
}

type JwtCustomClaims struct {
	UserID      int64  `json:"user_id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	RoleID      int32  `json:"role_id"`
	RoleName    string `json:"role_name"`
	Email       string `json:"email"`
	SectionID   int32  `json:"section_id"`
	SectionName string `json:"section_name"`
	jwt.StandardClaims
}
