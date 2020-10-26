package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println(`
Gauss is a cli for basic arithimetics.
	
Usage:

        gauss <command> [-f float] [arguments]

The commands are:

        add     gives the somatory of the given sequence
        sub     gives the somatory of the negative form of the given sequence
        prod    gives the productory of the given sequence
    `)
}

func main() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			printHelp()
		} else {

		}
	} else {
		fmt.Fprintln(os.Stderr, "Insufficient arguments...")
	}
}
