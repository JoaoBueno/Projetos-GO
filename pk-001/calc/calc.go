package calc

import "errors"

func Dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("denominador não pode ser 0")
	}
	return a / b, nil
}
