// +build echo_join_main

package main

import (
	"fmt";
	"os";
	"strings";
	"github.com/riccardotonini/time_monitoring";
)

func echo (args []string) {
	p := fmt.Println
	p(strings.Join(args, " "))
}

func main() {
    p := fmt.Println
    timed_echo := time_monitoring.Timed(echo)
    timed_echo(echo, os.Args[1:])
}
