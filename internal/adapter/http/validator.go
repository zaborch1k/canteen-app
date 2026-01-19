package http

import (
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Struct(v any) error
}

type playgroundValidator struct {
	v *validator.Validate
}

func NewValidator() Validator {
	v := validator.New()
	return &playgroundValidator{v: v}
}

func (p *playgroundValidator) Struct(v any) error {
	return p.v.Struct(v)
}
