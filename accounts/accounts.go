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

// Deposit x amount on your account
func (a *Account) Deposit(amount int) { // receiver, 첫 글자는 Struct의 첫 글자
	// *를 해주지 않으면 복사복의 Account를 받아온다. * -> pointer receiver
	a.balance += amount // go에서는 Class(struct)에서 메서드를 선언하지 않고 이런식으로 함.
}

// Balance of  your account's balance
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) {
	a.balance -= amount
}
