package accounts

// Account Struct
type Account struct { // export Upper -> public
	// Owner   string // public Upper
	owner   string // private Upper
	balance int
}

// NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
