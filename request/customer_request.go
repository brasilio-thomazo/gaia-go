package request

import "optimus.dev.br/gaia/model"

type CustomerRequest struct {
	Name     string          `json:"name"`
	Phone    string          `json:"phone"`
	Email    string          `json:"email"`
	Document string          `json:"document"`
	Address  string          `json:"address"`
	Contacts []model.Contact `json:"contacts"`
	Active   bool            `json:"active"`
}
