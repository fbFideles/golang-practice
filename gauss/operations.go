package gauss

import "errors"

type Operation struct {
	values []interface{}
}

func (o *Operation) Add(flag string) (sum interface{}, err error) {
	if len(o.values) <= 1 {
		return nil, errors.New("It's not possible to sum less than 2 operands...")
	}

	sumTemp := int64(0)
	for _, num := range o.values {
		sumTemp += num.(int64)
	}

	sum = sumTemp

	return
}
