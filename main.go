package main

import (
	"os"
	"fmt"
	"bufio"
	"github.com/geremachek/arzhur/frame"
)

func main() {
	var windows []string

	if len(os.Args) > 1 { // read from arguments, if they are supplied.
		windows = os.Args[1:]
	} else { // we are reading from stdin
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			windows = append(windows, scanner.Text()) // treat each line of input as its own window
		}

		// if nothing came in through stdin, supply a conciliatory empty string

		if len(windows) == 0 {
			windows = append(windows, "")
		}
	}

	// start the UI

	if f, err := frame.NewFrame(windows); err == nil {
		if out, err := f.Start(); err == nil {
			fmt.Print(out) // print the text of the selected window(s)
		} else {
			printError("couldn't start interface")
		}
	} else {
		printError("could not open screen")
	}
}

// print an error message and return 1

func printError(msg string) {
	fmt.Fprintf(os.Stderr, "arzhur: %s\n", msg)
	os.Exit(1)
}