// entry point, for compiling main.go is needed
package main

import (
	"fmt"
	"time"
)

// where saving the results
func main(){
	c := make(chan string)
	people := [5]string{"nico","flynn","dal","japanguy","jarry"}
	// create goroutines
	for _, person := range people {
		go isSexy(person, c)
	}
	for i:=0;i<len(people);i++{
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	// sleep 5 sec
	time.Sleep(time.Second*5)
	// send channel this message
	c <- person + " is sexy"
}