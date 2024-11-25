package response

import "optimus.dev.br/gaia/model"

type AuthResponse struct {
	model.User
	Token string `json:"token"`
}
