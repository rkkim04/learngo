// entry point, for compiling main.go is needed
package main

import (
	"fmt"
	"time"
)

// where saving the results
func main(){
	c := make(chan string)
	people := [2]string{"nico","flynn"}
	// create goroutines
	for _, person := range people {
		go isSexy(person, c)
	}
	resultOne := <-c
	resultTwo := <-c
	resultThree := <-c
	fmt.Println("Waiting for messages.")
	// blocking operations. 끝날때까지 멈춤 여기서
	// once c get a message then move to next line
	fmt.Println("Received this message:", resultOne)
	// and then moved to next and wait for the return
	fmt.Println("Received this message:", resultTwo)
	// and main func ends

	// goroutines 는 갯수만큼만 동작하므로 더 많이 입력하면 에러가 발생함
	fmt.Println("Received this message:", resultThree)
}

func isSexy(person string, c chan string) {
	// sleep 5 sec
	time.Sleep(time.Second*5)
	// send channel this message
	c <- person + " is sexy"
}