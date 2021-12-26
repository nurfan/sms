package model

type Findings struct {
	ID                 int64        `db:"id" json:"id"`
	Code               string       `db:"code" json:"code"`
	Status             int32        `db:"status" json:"status"`
	Date               string       `db:"date" json:"date"`
	SectionID          int32        `db:"section_id" json:"section_id"`
	Area               string       `db:"area" json:"area"`
	Pic                string       `db:"pic" json:"pic"`
	Description        string       `db:"description" json:"description"`
	AccidentID         int32        `db:"accident_id" json:"accident_id"`
	FindingsCategoryID int32        `db:"findings_category_id" json:"findings_category_id"`
	UnsafeCategoryID   int32        `db:"unsafe_category_id" json:"unsafe_category_id"`
	RiskCategoryID     int32        `db:"risk_category_id" json:"risk_category_id"`
	Object             string       `db:"object" json:"object"`
	Notes              string       `db:"notes" json:"notes"`
	Attachment         []Attachment `db:"attachment" json:"attachment"`
	CreatedBy          string       `db:"created_by" json:"created_by"`
	CreatedAt          string       `db:"created_at" json:"created_at"`
	UpdatedBy          string       `db:"updated_by" json:"updated_by"`
	UpdatedAt          string       `db:"updated_at" json:"updated_at"`
}

type Attachment struct {
	Type string `json:"type"`
	Path string `json:"path"`
}
