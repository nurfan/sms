package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/nurfan/sms/model"
	"github.com/pkg/errors"
)

// GetFindings ...
func (r *RepositoryPsql) FetchFindings(ctx context.Context) (result []model.Findings, err error) {

	query := `SELECT
					*
				FROM 
					findings f
				ORDER BY 
					status, date DESC  `

	err = r.Conn.SelectContext(ctx, &result, query)

	if err != nil && err != sql.ErrNoRows {
		return result, errors.Wrap(err, "repo.GetFindings")
	}

	return
}

// CreateFindings ...
func (r *RepositoryPsql) CreateFindings(ctx context.Context, params model.Findings) (model.Findings, error) {

	tx := r.Conn.MustBegin()
	var lastInsertId int
	err := tx.QueryRowx(`
		INSERT INTO addresses (
			code,
			status,
			date,
			section_id,
			area,
			pic,
			description,
			accident_id,
			findings_category_id,
			unsafe_category,
			risk_category,
			object,
			notes,
			attachment,
			created_by,
			updated_by
		) VALUES 
			($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18) 
		RETURNING id`,
		params.Code,
		params.Status,
		params.Date,
		params.SectionID,
		params.Area,
		params.Pic,
		params.Description,
		params.AccidentID,
		params.FindingsCategoryID,
		params.UnsafeCategoryID,
		params.RiskCategoryID,
		params.Object,
		params.Notes,
		params.Attachment,
		params.UpdatedBy,
		params.CreatedBy).Scan(&lastInsertId)
	if err != nil {
		tx.Rollback()
		return params, errors.Wrap(err, "insert address error")
	}
	log.Println("lastInsertId: ", lastInsertId)

	err = tx.Commit()
	if err != nil {
		return params, errors.Wrap(err, "tx.Commit()")
	}

	params.ID = int64(lastInsertId)
	return params, nil
}
