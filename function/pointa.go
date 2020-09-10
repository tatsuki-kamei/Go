package main

func swap2(x, y *int)(*int, *int){
	*x,*y = *y,*x
	return x,y
}


func main() {
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}
