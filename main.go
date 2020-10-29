package main

import (
	"fmt"
	"log"
	"os"

	"github.com/calculator-cli/gauss"
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

	var operation gauss.Operation
	operation.Values = make([]interface{}, 0)

	if len(os.Args) >= 2 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			printHelp()
		} else {
			switch {
			case os.Args[1] == "add":
				for i := 2; i < len(os.Args); i++ {
					operation.Values = append(operation.Values, os.Args[i])
				}
				sum, err := operation.Add(false)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(sum)
			}
		}
	} else {
		fmt.Fprintln(os.Stderr, "Insufficient arguments...")
	}
}
