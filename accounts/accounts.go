package accounts

import (
	"errors"
	"fmt"
)

// Account Struct
type Account struct {
	owner   string
	balance int
}

// error 이름은 err000 이어야 한다
var errNoMoney = errors.New("Can't withdraw you are poor")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // 새로운 오브젝트 전달
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	// (a Account) receiver -> struct의 첫 글자를 따서 소문자로 만들어야 함
	// (a *Account) -> account를 복사하지말고 Deposit method를 호출한 account를 사용하라는 의미
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
