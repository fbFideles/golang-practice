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

	if len(o.Values) <= 1 {
		return nil, errors.New("It's not possible to sum less than 2 operands")
	}

	if flag {
		sumTemp := int64(0)
		for _, num := range o.Values {

			var s string = num.(string)
			number, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			sumTemp += number

		}

		sum = sumTemp
	} else {
		sumTemp := float64(0)
		for _, num := range o.Values {
			var s string = num.(string)
			number, err := strconv.ParseFloat(s, 64)
			if err != nil {
				log.Fatal(err)
			}

			sumTemp += number
		}

		sum = sumTemp
	}

	return
}

// Sub -
func (o *Operation) Sub(flag bool) (sub interface{}, err error) {
	if len(o.Values) <= 1 {
		return nil, errors.New("It's not possible to subtract less than 2 operands")
	}

	if flag {
		subTemp := int64(0)
		for _, num := range o.Values {
			subTemp -= num.(int64)
		}

		sub = subTemp
	} else {
		subTemp := float64(0)
		for _, num := range o.Values {
			subTemp -= num.(float64)
		}

		sub = subTemp
	}

	return
}
