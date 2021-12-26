package action

import (
	"context"

	"github.com/nurfan/sms/model"
	repo "github.com/nurfan/sms/repository"
	"github.com/nurfan/sms/util/errors"
)

type FecthAccidents struct {
	e    errors.UniError
	repo *repo.RepositoryPsql
}

func (gt *FecthAccidents) Handle(ctx context.Context) ([]model.FormSelect, *errors.UniError) {

	result, err := gt.repo.FetchAccidents(ctx)
	if err != nil {
		return result, gt.e.SystemError(err)
	}

	return result, nil
}

func NewFecthAccidents(r *repo.RepositoryPsql) *FecthAccidents {
	return &FecthAccidents{
		repo: r,
	}
}
