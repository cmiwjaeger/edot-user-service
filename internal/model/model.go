package model

// WebResponse is a generic response wrapper
// @Description WebResponse is a generic response wrapper
// @Success 200 {object} WebResponse
type WebResponse[T any] struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    T             `json:"data"`
	Paging  *PageMetadata `json:"paging,omitempty"`
	Errors  string        `json:"errors,omitempty"`
}

type QueryListRequest struct {
	Keyword string `json:"keyword"`
	Page    int    `json:"page" validate:"required"`
	Size    int    `json:"size" validate:"required"`
	Sort    string `json:"sort"`
}

type PageResponse[T any] struct {
	Status       string       `json:"status"`
	Message      string       `json:"message"`
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"paging,omitempty"`
	Errors       string       `json:"errors,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"totalItem"`
	TotalPage int64 `json:"totalPage"`
}
