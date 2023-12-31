package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

type singleton struct {
}

var singletonInstance *singleton

func GetInstance() *singleton {
	if singletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singletonInstance == nil {
			fmt.Println("Creating single instance now.")
			singletonInstance = &singleton{}
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return singletonInstance
}
