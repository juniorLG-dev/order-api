package vo

import (
	"errors"
	"regexp"
)

type CEP struct {
	Value string `json:"code"`
}

func NewCEP(cep string) (*CEP, error) {
	match, _ := regexp.MatchString("^[0-9]{5}-[0-9]{3}$", cep)
	if !match {
		return nil, errors.New("invalid CEP")
	}
	return &CEP{
		Value: cep,
	}, nil
}
