package entities

import (
	"github.com/go-playground/validator/v10"
	"github.com/pborman/uuid"
)

type Tweet struct {
	ID          string `json:"id" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func NewTweet() *Tweet {
	tweet := Tweet{
		ID: uuid.New(),
	}

	return &tweet
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}
