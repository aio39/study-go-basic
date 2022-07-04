package main

import (
	"fmt"

	"github.com/aio39/study_go_basic/accounts"
)

func main() {
	// account := banking.Account{Owner: "aio", Balance: 1000}
	account := accounts.NewAccount("aio") // construct가 없는 go에서 private 프로퍼티를 초기화 하는 패턴
	account.Deposit(100)
	fmt.Println(account)
	fmt.Println(account.Balance())
}
