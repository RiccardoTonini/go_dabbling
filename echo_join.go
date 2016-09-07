// +build echo_join_main

package main

import (
	"fmt";
	"os";
	"strings";
)

func main() {
         fmt.Println(strings.Join(os.Args[1:], " "))
}
