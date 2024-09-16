package main

func main() {
	x := allocate()
	println(*x * 2)
}

func allocate() *int {
	x := 20
	return &x
}
