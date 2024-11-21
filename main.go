// entry point, for compiling main.go is needed
package main

import (
	"fmt"

	"github.com/rkkim04/learngo/banking"
)

func main(){
	account := banking.Account{Owner:"nicolas", Balance:1000}
	fmt.Println(account)
}

