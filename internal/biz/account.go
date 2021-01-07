package biz

// Account
type Account struct {
	Id int
	Name string
	Address Address
}

// Address
type Address struct {
	Id int
	City string
	Area string
}

// AccountRepo
type AccountRepo interface {
	SaveAccount(*Account)
}

// NewAccountUserCase
func NewAccountUserCase(repo AccountRepo) *AccountUserCase {
	return &AccountUserCase{repo: repo}
}

// AccountUserCase
type AccountUserCase struct {
	repo AccountRepo
}

// SaveAccount
func (uc *AccountUserCase) SaveAccount(a *Account) {
	uc.repo.SaveAccount(a)
}
