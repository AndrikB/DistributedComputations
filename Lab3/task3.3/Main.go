package main

import (
	"fmt"
	"math/rand"
	"time"
)

const SmokersNum = 3

func smoker(component int, semaphoreForSmoking chan bool, semaphoreForDealing chan bool, table *[]bool) {
	//var state = 0
	for {
		//state = 0
		<-semaphoreForSmoking
		fmt.Println("Checking table # ", component, "...")
		if !(*table)[component] {

			fmt.Println("Smoking # ", component, "...")
			time.Sleep(time.Second)
			for i := range *table {
				(*table)[i] = false
			}
			semaphoreForDealing <- true
			//state = 1
		}
		/*if state == 0 {
			semaphoreForSmoking <- true
		}*/
	}
}

func middle(semaphoreForSmoking chan bool, semaphoreForDealing chan bool, table *[]bool) {
	for {
		<-semaphoreForDealing
		fmt.Println("New items")
		var first, second = getCigaretteStuff()
		(*table)[first] = true
		(*table)[second] = true
		for i := 0; i < SmokersNum; i++ {
			semaphoreForSmoking <- true
		}
	}
}

func getCigaretteStuff() (int, int) {
	stuff1 := rand.Intn(3)
	stuff2 := rand.Intn(3)
	for stuff2 == stuff1 {
		stuff2 = rand.Intn(3)
	}

	return stuff1, stuff2
}

func main() {

	var table = make([]bool, SmokersNum, SmokersNum)
	var semaphoreForSmoking = make(chan bool, 3)
	var semaphoreForDealing = make(chan bool, 1)
	semaphoreForDealing <- true
	go middle(semaphoreForSmoking, semaphoreForDealing, &table)
	for i := 0; i < SmokersNum; i++ {
		go smoker(i, semaphoreForSmoking, semaphoreForDealing, &table)
	}
	_, _ = fmt.Scanln()
}
