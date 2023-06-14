package main

import "time"

func main() {

	for i := 0; i < 30; i++ {
		go GetInstance()
	}

	time.Sleep(10 * time.Second)

}
