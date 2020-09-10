package main

func kusucheck(x int){
	switch x%2{
	case 0:
		println("-偶数")
	case 1:
		println("-奇数")
	}
}

func main(){
	kusucheck(4)
}
