package action

import (
	"context"

	"github.com/nurfan/sms/model"
	repo "github.com/nurfan/sms/repository"
	"github.com/nurfan/sms/util/errors"
)

type CreateFindings struct {
	e    errors.UniError
	repo *repo.RepositoryPsql
}

func (cf *CreateFindings) Handle(ctx context.Context, req model.Findings) (model.Findings, *errors.UniError) {
	var result model.Findings

	findings, err := cf.repo.CreateFindings(ctx, req)
	if err != nil {
		return result, cf.e.SystemError(err)
	}

	return findings, nil
}

func NewCreateFindings(r *repo.RepositoryPsql) *CreateFindings {
	return &CreateFindings{
		repo: r,
	}
}
