package main

import (
	"fmt"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	reciientChannel := make(chan Recipient) // unbuffered channel
	go func() {
		err := loadRecipient("./users_100.csv", reciientChannel)
		if err != nil {
			fmt.Println(err)
		}
	}()
	var wg sync.WaitGroup

	workerCount := 5
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, reciientChannel, &wg)
	}
	wg.Wait()
}
