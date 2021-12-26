package repository

import (
	"context"
	"database/sql"

	"github.com/nurfan/sms/model"
	"github.com/pkg/errors"
)

// FetchAccidents ...
func (r *RepositoryPsql) FetchAccidents(ctx context.Context) (result []model.FormSelect, err error) {

	query := `SELECT
					id
				FROM 
					accidents ac
				WHERE 
					deleted_at IS NULL
				ORDER BY 
					name ASC`

	err = r.Conn.SelectContext(ctx, &result, query)

	if err != nil && err != sql.ErrNoRows {
		return result, errors.Wrap(err, "repo.FetchAccidents")
	}

	return
}

// FetchObjectCategory ...
func (r *RepositoryPsql) FetchObjectCategory(ctx context.Context) (result []model.FormSelect, err error) {

	query := `SELECT
					id
				FROM 
					object_category oc
				WHERE 
					deleted_at IS NULL
				ORDER BY 
					name ASC`

	err = r.Conn.SelectContext(ctx, &result, query)

	if err != nil && err != sql.ErrNoRows {
		return result, errors.Wrap(err, "repo.FetchObjectCategory")
	}

	return
}
