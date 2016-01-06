package main

import (
	"fmt"
	"time"
)

// START OMIT
type Ball struct{ hits int }

var i int

func main() {
    table := make(chan *Ball)
    go player("ping", table)
    go player("pong", table)

    table <- new(Ball) // game on; toss the ball
	i++
    time.Sleep(1 * time.Second)
    <-table
}

func player(name string, table chan *Ball) {
    for {
		i++
        ball := <-table
        ball.hits++
        fmt.Println(name, ball.hits)
        time.Sleep(100 * time.Millisecond)
        table <- ball
    }
}
// END OMIT
