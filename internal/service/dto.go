package service

type PaginationParams struct {
	Page  int32 `json:"page" form:"page" binding:"required,numeric,gt=0"`
	Limit int32 `json:"limit" form:"limit" binding:"required,numeric,gt=0"`
}