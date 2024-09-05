package main

import (
	"fmt"
	"time"

	"github.com/dejangegic/gravedigger/cli"
)

func main() {
	timer := time.Now()
	cli.Execute()
	fmt.Println("Elapsed time: " + time.Since(timer).String())

}
