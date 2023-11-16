package model

type PagingQuery struct {
	Page  int `form:"page" validate:"required,min=1"`
	Limit int `form:"limit" validate:"required,min=1"`
}
