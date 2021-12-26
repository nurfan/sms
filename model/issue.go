package model

type GetFindingsRequest struct {
	UserID    int32 `json:"user_id"`
	RoleID    int32 `json:"role_id"`
	SectionID int32 `json:"section_id"`
	Page      int32 `json:"page"`
	Limit     int32 `json:"limit"`
}

type GetFindingsResponse struct {
}
