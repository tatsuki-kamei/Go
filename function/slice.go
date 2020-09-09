package main

import "fmt"

func main(){
	n := []int{19,86,1,12}
	var sum int
	for i := 0;i < len(n); i +=1{
		sum += n[i]
	}
	fmt.Print(sum)
}