package main

import (
	"time"
	"api/mux"
)

func main(){
	go mux.Run()

	for {
		time.Sleep(5 * time.Second)
	}

}