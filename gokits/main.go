package main

func main() {
	println(f(1))
}

func f(x int) (_, __ int) {
	_, __ = x, x
	return
}
