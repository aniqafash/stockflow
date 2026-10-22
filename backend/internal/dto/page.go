package dto

const (
	BaseLimit int = 10
	MaxLimit  int = 50
	BasePage  int = 1
)

type Page struct {
	Page  *int `form:"page"`
	Limit *int `form:"limit"`
}
