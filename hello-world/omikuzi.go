package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6)
	switch s + 1{
	case  6:
		fmt.Printf("大吉")
	case  5,4:
		fmt.Printf("中吉")
	case  3,2:
		fmt.Printf("吉")
	default:
		fmt.Printf("凶")
	}
}