package http

import (
	"github.com/nurfan/sms/repository"
	"github.com/nurfan/sms/util/errors"
)

type Adapter struct {
	repoPsql *repository.RepositoryPsql
	e        errors.UniError
}

func NewAdapter(repoPsql *repository.RepositoryPsql) *Adapter {
	return &Adapter{
		repoPsql: repoPsql,
	}
}
