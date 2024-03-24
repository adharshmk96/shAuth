package core

import "github.com/adharshmk96/shAuth/core/model"

type AccountStorage interface {
	SaveAccount(account *model.Account) error
	GetAccountByEmail(email string) (*model.Account, error)
	UpdateAccount(account *model.Account) error
}
