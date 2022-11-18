package book

import "encoding/json"

type BookRequest struct {
	// https://pkg.go.dev/github.com/go-playground/validator/v10#readme-comparisons
	Title    string      `json:"title" binding:"required"`
	Price    json.Number `json:"price" binding:"required,numeric"`
	Subtitle string      `json:"subtitle"`
}
