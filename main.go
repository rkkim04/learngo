// entry point, for compiling main.go is needed
package main

import (
	"fmt"

	"github.com/rkkim04/learngo/mydict"
)

func main(){
	// account := accounts.NewAccount("lois")
	// account.Deposit(1000)
	// err := account.Withdraw(2000)
	// if err != nil{
	// 	log.Println(err)
	// }
	// fmt.Println(account)

	dictionary := mydict.Dictionary{}
	// definition, err := dictionary.Search("first")
	// if err != nil{
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Println(definition)
	// }
	// word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// Search()로 Dictionary에 word가 있는지 확인
	// hello, _ := dictionary.Search(word)
	// fmt.Println("found", word, "definition:", hello)

	// 단어추가: 이미 있는 단어인 경우 미리 지정한 err 메세지 출력
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
	fmt.Println(word)
}

