package services

import (
	"github.com/fanda-org/postmasters/config"
	"github.com/fanda-org/postmasters/database"
)

// NewUsersService creates UsersService
func NewUsersService(dbConfig *config.DBConfig) *UsersService {
	service := UsersService{dbConfig: dbConfig}
	service.db = database.Open(dbConfig)
	return &service
}
