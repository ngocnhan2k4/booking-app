package main

import ("fmt"
		"time"
		"sync"
)

func sendTicket(userName string){
	time.Sleep(10 * time.Second)
	fmt.Printf("Send ticket %v", userName)
	wg.Done() // Decrease counter by 1
}

var wg = sync.WaitGroup{}
func main(){
	wg.Add(2) // Add 2 goroutines to wait group
	go sendTicket("Nhan")
	fmt.Println("Main function")
	go sendTicket("Hai")

	wg.Wait() // block cho đến khi counter = 0
}