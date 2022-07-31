package core

import (
	"math/big"

	"github.com/numary/ledger/pkg/core/monetary"
)

type MonetaryInt = *big.Int

type Monetary struct {
	Asset  string       `json:"asset"`
	Amount monetary.Int `json:"amount"`
}

var MonetaryIntZero = monetary.NewInt(0)

func NewMonetaryInt(amount int64) *monetary.Int {
	return monetary.NewInt(amount)
}

func ParseMonetaryInt(amount string) *monetary.Int {
	return monetary.NewIntFromString(amount)
}

func AddMonetaryInt(a, b *monetary.Int) *monetary.Int {
	return a.Add(b)
}

func SubMonetaryInt(a, b *monetary.Int) *monetary.Int {
	return a.Sub(b)
}

func NegMonetaryInt(a *monetary.Int) *monetary.Int {
	return a.Neg()
}

// todo: fix this
func LtMonetaryInt(a, b *monetary.Int) bool {
	return a.Lt(b)
}

// todo: fix this
func LteMonetaryInt(a, b *monetary.Int) bool {
	return a.Lte(b)
}

// todo: fix this
func GteMonetaryInt(a, b *monetary.Int) bool {
	return a.Gte(b)
}

// todo: fix this
func GtMonetaryInt(a, b *monetary.Int) bool {
	return a.Gt(b)
}

func EqMonetaryInt(a, b *monetary.Int) bool {
	return a.Eq(b)
}
