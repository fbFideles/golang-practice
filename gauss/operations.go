package gauss

import "errors"

type Operation struct {
	values []interface{}
}

func (o *Operation) Add(flag bool) (sum interface{}, err error) {
	if len(o.values) <= 1 {
		return nil, errors.New("It's not possible to sum less than 2 operands...")
	}

	if flag {
		sumTemp := int64(0)
		for _, num := range o.values {
			sumTemp += num.(int64)
		}

		sum = sumTemp
	} else {
		sumTemp := float64(0)
		for _, num := range o.values {
			sumTemp += num.(float64)
		}

		sum = sumTemp
	}

	return
}

func (o *Operation) Sub(flag bool) (sub interface{}, err error) {
	if len(o.values) <= 1 {
		return nil, errors.New("It's not possible to subtract less than 2 operands...")
	}

	if flag {
		subTemp := int64(0)
		for _, num := range o.values {
			subTemp -= num.(int64)
		}

		sub = subTemp
	} else {
		subTemp := float64(0)
		for _, num := range o.values {
			subTemp -= num.(float64)
		}

		sub = subTemp
	}

	return
}
