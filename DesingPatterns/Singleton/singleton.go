package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Create new instance")
			singleInstance = &single{}
		} else {
			fmt.Println("single instance already created-1 with locked")
		}
	} else {
		fmt.Println("single instance already created-2")
	}
	return singleInstance
}

func getInstanceWithSyncOnce() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
