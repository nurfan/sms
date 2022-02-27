package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/nurfan/sms/model"
)

// GetReminderLog ...
func (r *RepositoryPsql) GetAccount(ctx context.Context, username string) (result model.Account, err error) {

	query := `SELECT
					acc.user_id,
					acc.username,
					acc.nickname,
					acc.password,
					acc.email,
					acc.section_id,
					acc.role_id,
					ro.name as role_name,
					acc.is_active,
					acc.last_login,
					acc.created_by,
					acc.created_at,
					acc.updated_by,
					acc.updated_at,
					acc.section_id,
					se.name AS section_name
				FROM 
					accounts acc
				JOIN 
					roles ro ON acc.role_id = ro.id 
					sections se ON se.id = acc.section_id
				WHERE
					acc.username = $1`

	err = r.Conn.GetContext(ctx, &result, query, username)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("repo.GetAccount : %s", err)
		return
	}

	return
}
