package main

func main() {
	c := make(chan int)
	c <- 42
	val := <-c
	println(val)
}
