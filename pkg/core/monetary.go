package core

import (
	"math/big"
)

type MonetaryInt = *big.Int

type Monetary struct {
	Asset  string      `json:"asset"`
	Amount MonetaryInt `json:"amount"`
}

var MonetaryIntZero = NewMonetaryInt(0)

func NewMonetaryInt(amount int64) MonetaryInt {
	return MonetaryInt(big.NewInt(amount))
}

func ParseMonetaryInt(amount string) MonetaryInt {
	i, _ := big.NewInt(0).SetString(amount, 10)
	return MonetaryInt(i)
}

func AddMonetaryInt(a, b MonetaryInt) MonetaryInt {
	if a == nil {
		a = big.NewInt(0)
	}

	if b == nil {
		b = big.NewInt(0)
	}

	return MonetaryInt(big.NewInt(0).Add(a, b))
}

func SubMonetaryInt(a, b MonetaryInt) MonetaryInt {
	if a == nil {
		a = big.NewInt(0)
	}

	if b == nil {
		b = big.NewInt(0)
	}

	return MonetaryInt(big.NewInt(0).Sub(a, b))
}

func NegMonetaryInt(a MonetaryInt) MonetaryInt {
	return MonetaryInt(big.NewInt(0).Neg(a))
}

// todo: fix this
func LtMonetaryInt(a, b MonetaryInt) bool {
	return big.NewInt(0).Cmp(a) < 0
}

// todo: fix this
func LteMonetaryInt(a, b MonetaryInt) bool {
	return a.Cmp(b) <= 0
}

// todo: fix this
func GteMonetaryInt(a, b MonetaryInt) bool {
	return a.Cmp(b) >= 0
}

// todo: fix this
func GtMonetaryInt(a, b MonetaryInt) bool {
	return a.Cmp(b) > 0
}

func EqMonetaryInt(a, b MonetaryInt) bool {
	return a.Cmp(b) == 0
}
