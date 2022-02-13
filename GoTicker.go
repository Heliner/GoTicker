package main

import (
	"fmt"
	"os"
	"strconv"
)

import (
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	if len(os.Args) < 1 {
		fmt.Println("empty args")
		return
	}
	sleep_minues, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("not a int")
		return
	}

	tick := time.Tick(1 * time.Minute)
	for ; sleep_minues > 0; sleep_minues-- {
		fmt.Printf("\r%2d", sleep_minues)
		<-tick
	}
	fmt.Println("\rfinish!")
}
