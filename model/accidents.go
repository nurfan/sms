package model

type FormSelect struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	CreatedBy string `db:"created_by"`
	CreatedAt string `db:"created_at"`
	UpdatedBy string `db:"updated_by"`
	UpdatedAt string `db:"updated_at"`
	DeletedBy string `db:"deleted_by"`
	DeletedAt string `db:"deleted_at"`
}
