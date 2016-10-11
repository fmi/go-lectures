package main

func main() {
	number := 42
	specificNumber := float64(number)
	println(number * specificNumber) // (mismatched types int and float64)
}
