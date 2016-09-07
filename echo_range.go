
package main

import (
	"fmt";
	"os";
	"strconv";
	"reflect";
	"github.com/riccardotonini/time_monitoring";
)

func echo_with_range(args []string) {
	s, sep := "", ""
	p := fmt.Println
	p("TYPE OF args IS ", reflect.TypeOf(args))
	for index, arg := range args {
		// p("THE NAME OF THE COMMAND IS: " + os.Args[0])
		p("index " + strconv.Itoa(index) + " arg " + arg)
		// p(s)
		s += sep + arg
		sep = " "
	}
	p(s)
}

func main() {
	//s, sep := "", "" // Alternatives: var s, sep string or var s = "", var sep = ""
	// range produces a pair of values: the index
	// and the value of the element at that index.
	// The blank identifier used whenever syntax requires a variable name
	// but program logic does not.

	timed_echo_with_range := time_monitoring.Timed(echo_with_range)
	timed_echo_with_range((os.Args[1:]))
}