package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
	"strings"
	"github.com/geremachek/arzhur/frame"
)

func main() {
	empty := flag.Bool("n", false, "Open an empty window on start.")
	noSplit := flag.Bool("s", false, "Don't split input into lines.")
	flag.Parse()

	var windows []string

	if *empty {
		windows = []string{""} // empty window
	} else if args := flag.Args(); len(args) >= 1 { // read from arguments, if they are supplied.
		windows = args
	} else { // we are reading from stdin
		scanner := bufio.NewScanner(os.Stdin)

		if *noSplit { // join input by newlines
			builder := strings.Builder{}

			for scanner.Scan() {
				builder.WriteString(scanner.Text() + "\n")
			}

			windows = append(windows, builder.String()) // append the input as a single window
		} else {
			for scanner.Scan() {
				windows = append(windows, scanner.Text()) // treat each line of input as its own window
			}
		}

		// if nothing came in through stdin, supply a conciliatory empty string

		if len(windows) == 0 {
			windows = []string{""}
		}
	}

	// start the UI

	if f, err := frame.NewFrame(windows); err == nil {
		if out, err := f.Start(); err == nil {
			fmt.Print(out) // print the text of the selected window(s)
		} else {
			printError(err)
		}
	} else {
		printError(err)
	}
}

// print an error message

func printError(err error) {
	fmt.Fprintf(os.Stderr, "arzhur: %s\n", err)
	os.Exit(1)
}