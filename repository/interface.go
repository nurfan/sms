package repository

import (
	"context"

	"github.com/nurfan/sms/model"
)

type Psql interface {
	GetUser(ctx context.Context, username string) (result model.Account, err error)
	FetchFindings(ctx context.Context) (result []model.Findings, err error)
	CreateFindings(ctx context.Context, params model.Findings) (model.Findings, error)
	FetchAccidents(ctx context.Context) (result []model.FormSelect, err error)
	FetchObjectCategory(ctx context.Context) (result []model.FormSelect, err error)
}
