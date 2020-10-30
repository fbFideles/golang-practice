package gauss

import (
	"errors"
	"log"
	"strconv"
)

// Operation -
type Operation struct {
	Values []interface{}
}

// Add -
func (o *Operation) Add(flag bool) (sum interface{}, err error) {
	var (
		intSum   int64
		floatSum float64
	)

	if len(o.Values) <= 1 {
		return nil, errors.New("It's not possible to sum less than 2 operands")
	}

	for _, num := range o.Values {

		s, ok := num.(string)
		if !ok {
			return nil, errors.New("Type value error: not a string")
		}

		if flag {
			number, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			intSum += number
		} else {
			s := num.(string)
			number, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Fatal(err)
			}

			floatSum += number
		}
	}

	if flag {
		sum = intSum
	} else {
		sum = floatSum
	}

	return
}

// Sub -
func (o *Operation) Sub(flag bool) (sub interface{}, err error) {
	var (
		intSub   int64
		floatSub float64
		firt     bool = true
	)

	if len(o.Values) <= 1 {
		return nil, errors.New("It's not possible to subtract less than 2 operands")
	}

	for _, num := range o.Values {

		s, ok := num.(string)
		if !ok {
			return nil, errors.New("Type value error: not a string")
		}

		if flag {
			number, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			if firt {
				intSub = number
			} else {
				intSub -= number
			}

			firt = false

		} else {
			s := num.(string)
			number, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Fatal(err)
			}

			if firt {
				floatSub = number
			} else {
				floatSub -= number
			}

			firt = false
		}
	}

	if flag {
		sub = intSub
	} else {
		sub = floatSub
	}

	return
}

// Prod -
func (o *Operation) Prod(flag bool) (prod interface{}, err error) {
	var (
		intProd   int64   = 1
		floatProd float64 = 1
	)

	if len(o.Values) <= 1 {
		return nil, errors.New("It's not possible to productory less than 2 operands")
	}

	for _, num := range o.Values {

		s, ok := num.(string)
		if !ok {
			return nil, errors.New("Type value error: not a string")
		}

		if flag {
			number, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			intProd *= number
		} else {
			s := num.(string)
			number, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Fatal(err)
			}

			floatProd *= number
		}
	}

	if flag {
		prod = intProd
	} else {
		prod = floatProd
	}

	return
}
