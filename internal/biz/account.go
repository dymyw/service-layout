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
	SaveAccount(*Account) bool
	GetInfo(string) (*Account, error)
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
func (uc *AccountUserCase) SaveAccount(a *Account) bool {
	return uc.repo.SaveAccount(a)
}

// GetInfo
func (uc *AccountUserCase) GetInfo(id string) (*Account, error) {
	return uc.repo.GetInfo(id)
}
