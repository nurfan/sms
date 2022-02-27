package action

import (
	"context"
	"time"

	"github.com/nurfan/sms/model"
	repo "github.com/nurfan/sms/repository"
	"github.com/nurfan/sms/util/errors"

	"github.com/golang-jwt/jwt"
)

type GetToken struct {
	e    errors.UniError
	repo *repo.RepositoryPsql
}

func (gt *GetToken) Handle(ctx context.Context, req model.GetTokenRequest) (*model.GetTokenResponse, *errors.UniError) {
	var result model.GetTokenResponse

	// get account
	users, err := gt.repo.GetAccount(ctx, req.Username)
	if err != nil {
		return nil, gt.e.SystemError(err)
	}

	// Set custom claims
	claims := &model.JwtCustomClaims{
		users.UserID,
		users.Username,
		users.Nickname,
		users.RoleID,
		users.RoleName,
		users.Email,
		users.SectionID,
		users.SectionName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secretbanget"))
	if err != nil {
		return nil, gt.e.SystemError(err)
	}

	result.Data.Token = t
	result.Message = "successfully generate client token"

	return &result, nil
}

func NewGetToken(r *repo.RepositoryPsql) *GetToken {
	return &GetToken{
		repo: r,
	}
}
