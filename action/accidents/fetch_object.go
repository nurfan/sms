package action

import (
	"context"

	"github.com/nurfan/sms/model"
	repo "github.com/nurfan/sms/repository"
	"github.com/nurfan/sms/util/errors"
)

type FecthObject struct {
	e    errors.UniError
	repo *repo.RepositoryPsql
}

func (gt *FecthObject) Handle(ctx context.Context) ([]model.FormSelect, *errors.UniError) {

	result, err := gt.repo.FetchObjectCategory(ctx)
	if err != nil {
		return result, gt.e.SystemError(err)
	}

	return result, nil
}

func NewFecthObject(r *repo.RepositoryPsql) *FecthObject {
	return &FecthObject{
		repo: r,
	}
}
