package service

import (
	"github.com/adharshmk96/shAuth/core"
)

type Service struct {
	accountStorage core.AccountStorage
}

func New(accountStorage core.AccountStorage) core.AccountService {
	return &Service{accountStorage: accountStorage}
}
