package l1

import "math/big"

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

// Комментарий: в Go тип int справится с такими числами, но обратите внимание на возможное переполнение для ещё больших значений.
// Для очень больших чисел можно использовать math/big.

func AddBig(a, b *big.Int) *big.Int{
	var c big.Int
	c.Add(a, b)
	return &c
}

func SubBig(a, b *big.Int) *big.Int{
	var c big.Int
	c.Sub(a, b)
	return &c
}

func MulBig(a, b *big.Int) *big.Int{
	var c big.Int
	c.Mul(a, b)
	return &c
}

func DivBig(a, b *big.Int) *big.Float {
	var c big.Float
	floatA := new(big.Float).SetInt(a)
	floatB := new(big.Float).SetInt(b)
	c.Quo(floatA, floatB)
	return &c
}
