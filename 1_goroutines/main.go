package main

import (
	"fmt"
	"time"
)
func numbers() {  
	for i := 0; i < 5; i++ {
        time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ",i)
	}
}

func alphabets() {  
	for i := 'a'; i < 'f'; i++ {
        time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ",i)
	}
}

func main()  {
	go numbers()
	go alphabets()
    time.Sleep(3000 * time.Millisecond)
	fmt.Println("Main function!")
}
