package main

import (
	"fmt"

	"github.com/aio39/study_go_basic/accounts"
)

func main() {
	// account := banking.Account{Owner: "aio", Balance: 1000}
	account := accounts.NewAccount("aio") // construct가 없는 go에서 private 프로퍼티를 초기화 하는 패턴
	account.Deposit(100)
	err := account.Withdraw(200)
	if err != nil { // go에서는 에러 Exception을 캐치해주지 않아도 프로그램이 종료되지 않음.
		// err catch 보다는 귀찮지만, err 체크를 강제하므로 안정적임
		// log.Fatalln(err)
		fmt.Println(err)
	}

	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

}
