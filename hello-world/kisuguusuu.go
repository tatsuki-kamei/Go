package main

import "fmt"

func main(){
	for i :=0; i<= 100; i +=1{
		if i % 2 == 0{
			fmt.Printf("%d-偶数\n",i)
		}else{
			fmt.Printf("%d-奇数\n",i)
		}
	}

	for i :=0; i <= 100; i+=1{
		switch{
		case i %2 == 0:
			fmt.Printf("%d-偶数\n",i)
		default:
			fmt.Printf("%d-奇数\n",i)
		}
	}
}
