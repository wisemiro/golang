package services

import "auth/server/auth/models"
//service struct
type Service struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	User
}
