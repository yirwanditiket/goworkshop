package main

func NewStructCopy() MyStruct {
	return MyStruct{
		A: 123, B: 456, C: 789,
		D: "ABC", E: "DEF", F: "HIJ",
		G: true, H: true, I: true,
	}
}

func NewStructPointer() *MyStruct {
	return &MyStruct{
		A: 123, B: 456, C: 789,
		D: "ABC", E: "DEF", F: "HIJ",
		G: true, H: true, I: true,
	}
}

type MyStruct struct {
	A, B, C int
	D, E, F string
	G, H, I bool
}

func (s MyStruct) AddTwo(x int) int {
	return s.A + x + 2
}

func main() {
}


