package action

import (
	"context"

	"github.com/nurfan/sms/model"
	repo "github.com/nurfan/sms/repository"
	"github.com/nurfan/sms/util/errors"
)

type GetFindings struct {
	e    errors.UniError
	repo *repo.RepositoryPsql
}

func (gt *GetFindings) Handle(ctx context.Context, req model.GetFindingsRequest) ([]model.Findings, *errors.UniError) {
	var result []model.Findings

	findings, err := gt.repo.FetchFindings(ctx)
	if err != nil {
		return result, gt.e.SystemError(err)
	}

	return findings, nil
}

func NewGetFindings(r *repo.RepositoryPsql) *GetFindings {
	return &GetFindings{
		repo: r,
	}
}
