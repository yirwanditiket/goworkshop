package main

func main() {
	x := 4
	inc(&x)
	println(x)
}

func inc(x *int) {
	*x++
}
