// 패키지명을 꼭 붙어야함!
package main

import (
	"fmt"
	"golang-test/accounts"
	"golang-test/mydict"
)

// main package와 그 안에 있는 main function을 먼저 찾고 실행시킴
// main은 컴파을 위해 필요한것!
// terminal > go run main.go
func main() {
	/*
		// GOROOT 검색 후 GOPATH 검색하여 겨로를 사용
		song := testlib.GetMusic("Alicia Keys")
		fmt.Println(song)
		//fmt.Println(quote.Go())
		something.SayHello()
	*/

	//account := accounts.BankAccount{Owner: "kyo", Balance: 1000}
	account := accounts.NewAccount("kyo")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
		//log.Fatalln(err)
	}
	fmt.Println(account)

	dictionary := mydict.Dictionary{}
	//dictionary := mydict.Dictionary{"first": "First word"}
	//dictionary["hello"] = "hello"

	word := "hello"
	definition := "Greeting"
	dictionary.Add(word, definition)

	err = dictionary.Update("11", "Second")
	if err != nil {
		fmt.Println(err)
	}

	dictionary.Delete(word)

	_, err = dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dictionary)

}
