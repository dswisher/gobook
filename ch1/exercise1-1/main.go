// prints its command-line arguments, one per line with index

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%2d: %s\n", i, arg)
	}
}
