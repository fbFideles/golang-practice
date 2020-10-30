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
		deco	gives the prime decomposition of the given sequence
		dist    gives the two points distance
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
			case os.Args[1] == "sub":
				for i := 2; i < len(os.Args); i++ {
					operation.Values = append(operation.Values, os.Args[i])
				}
				sub, err := operation.Sub(false)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(sub)

			case os.Args[1] == "prod":
				for i := 2; i < len(os.Args); i++ {
					operation.Values = append(operation.Values, os.Args[i])
				}
				prod, err := operation.Prod(false)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(prod)

			case os.Args[1] == "deco":
				for i := 2; i < len(os.Args); i++ {
					operation.Values = append(operation.Values, os.Args[i])
				}
				primes, err := operation.DecomposePrimes()
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(primes)

			case os.Args[1] == "dist":
				for i := 2; i < len(os.Args); i++ {
					operation.Values = append(operation.Values, os.Args[i])
				}
				dist, err := operation.PointDistance()
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(*dist)

			}

			}
		}
	} else {
		fmt.Fprintln(os.Stderr, "Insufficient arguments...")
	}
}
