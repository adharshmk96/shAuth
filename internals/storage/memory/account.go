package memory

import (
	"fmt"
	"github.com/adharshmk96/shAuth/core"
	"github.com/adharshmk96/shAuth/core/model"
)

func errWrap(method string, err error) error {
	return fmt.Errorf("storage.%s: %w", method, err)
}

func (s *Storage) SaveAccount(account *model.Account) error {
	// if email already exists
	if _, err := s.GetAccountByEmail(account.Email); err == nil {
		return errWrap("SaveAccount", core.ErrAccountExists)
	}

	s.Accounts[account.ID.String()] = *account
	return nil
}

func (s *Storage) GetAccountByEmail(email string) (*model.Account, error) {
	acc := model.Account{}

	for _, a := range s.Accounts {
		if a.Email == email {
			acc = a
			return &acc, nil
		}
	}

	return nil, errWrap("GetAccountByEmail", core.ErrAccountNotFound)
}

func (s *Storage) UpdateAccount(account *model.Account) error {
	if _, ok := s.Accounts[account.ID.String()]; !ok {
		return errWrap("UpdateAccount", core.ErrAccountNotFound)
	}

	s.Accounts[account.ID.String()] = *account
	return nil
}
