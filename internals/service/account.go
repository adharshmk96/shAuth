package service

import (
	"fmt"
	"time"

	"github.com/adharshmk96/shAuth/core"
	"github.com/adharshmk96/shAuth/core/model"
	"github.com/adharshmk96/shAuth/pkg/utils"
	"github.com/google/uuid"
)

func errWrap(method string, err error) error {
	return fmt.Errorf("service.%s: %w", method, err)
}

func (s *Service) RegisterAccount(acc *model.Account) error {
	id := uuid.New()
	acc.ID = id

	password := acc.Password
	hash, err := utils.HashPassword(password)
	if err != nil {
		return errWrap("RegisterAccount", err)
	}
	acc.Password = hash

	acc.Created = time.Now()
	acc.Updated = time.Now()

	err = s.accountStorage.SaveAccount(acc)
	if err != nil {
		return errWrap("RegisterAccount", err)
	}

	return nil
}

func (s *Service) Authenticate(email, password string) (*model.Account, error) {
	acc, err := s.accountStorage.GetAccountByEmail(email)
	if err != nil {
		return acc, errWrap("Authenticate", err)
	}

	valid, err := utils.VerifyPassword(acc.Password, password)
	if !valid {
		return nil, errWrap("Authenticate", core.ErrInvalidCredentials)
	}
	if err != nil {
		return nil, errWrap("Authenticate", err)
	}

	return acc, nil
}

func (s *Service) GetAccountByEmail(email string) (*model.Account, error) {
	acc, err := s.accountStorage.GetAccountByEmail(email)
	if err != nil {
		return acc, errWrap("GetAccountByEmail", err)
	}

	return acc, nil
}

func (s *Service) ChangePassword(email, password, newPassword string) error {
	acc, err := s.Authenticate(email, password)
	if err != nil {
		return errWrap("ChangePassword", err)
	}

	hash, err := utils.HashPassword(newPassword)
	if err != nil {
		return errWrap("ChangePassword", err)
	}

	acc.Password = hash
	acc.Updated = time.Now()

	err = s.accountStorage.UpdateAccount(acc)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GenerateJWT(acc *model.Account) (string, error) {
	claims := utils.NewClaims()
	(*claims)["email"] = acc.Email

	tokenString, err := utils.EncodeJWT(claims)
	if err != nil {
		return "", errWrap("GenerateJWT", err)
	}

	return tokenString, nil
}

func (s *Service) ValidateJWT(token string) (*model.Account, error) {
	claims, err := utils.DecodeJWT(token)
	if err != nil {
		return nil, errWrap("ValidateJWT", err)
	}

	email := claims["email"].(string)

	acc, err := s.accountStorage.GetAccountByEmail(email)
	if err != nil {
		return nil, errWrap("ValidateJWT", err)
	}

	if acc.Email != email {
		return nil, errWrap("ValidateJWT", core.ErrInvalidCredentials)
	}

	return acc, nil
}
