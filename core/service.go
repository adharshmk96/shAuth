package core

import "github.com/adharshmk96/shAuth/core/model"

type AccountService interface {
	// Auth
	RegisterAccount(acc *model.Account) error
	Authenticate(email, password string) (*model.Account, error)
	GetAccountByEmail(email string) (*model.Account, error)
	ChangePassword(email, password, newPassword string) error

	// JWT
	GenerateJWT(account *model.Account) (string, error)
	ValidateJWT(token string) (*model.Account, error)
}
