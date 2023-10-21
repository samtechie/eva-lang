package eva

import (
	"fmt"
)

type Eva struct{}

func (e Eva) eval(exp interface{}) (interface{}, error) {
	if isNumber(exp) {
		return exp, nil
	}
	return nil, fmt.Errorf("Unimplemented")
}

func isNumber(exp interface{}) bool {
	_, ok := exp.(int)
	return ok
}
