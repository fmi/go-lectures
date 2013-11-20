package main

func fib() <-chan int {
	c := make(chan int)
	go func() {
		for a, b := 0, 1; ; a, b = b, a+b {
			c <- a
		}
	}()
	return c
}

func main() {
	fibonacci := fib()
	for i := 0; i < 10; i++ {
		println(<-fibonacci)
	}
}
