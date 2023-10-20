package models

type UserDataResponse struct { // TODO: pasar a models
	Data struct {
		Email string `json:"email"`
	}
}
