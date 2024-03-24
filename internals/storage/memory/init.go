package memory

import (
	"github.com/adharshmk96/shAuth/core"
	"github.com/adharshmk96/shAuth/core/model"
)

type Storage struct {
	Accounts map[string]model.Account
}

func New() core.AccountStorage {
	return &Storage{Accounts: make(map[string]model.Account)}
}
