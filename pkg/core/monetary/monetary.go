package monetary

import "math/big"

type Int big.Int

func NewInt(amount int64) *Int {
	return (*Int)(big.NewInt(amount))
}

func NewIntFromString(v string) *Int {
	i, _ := big.NewInt(0).SetString(v, 10)
	return (*Int)(i)
}

func (x *Int) Equal(y *Int) bool {
	return true
}

func (x *Int) Gte(y *Int) bool {
	return (*big.Int)(x).Cmp((*big.Int)(y)) >= 0
}

func (x *Int) Gt(y *Int) bool {
	return (*big.Int)(x).Cmp((*big.Int)(y)) > 0
}

func (x *Int) Lte(y *Int) bool {
	return (*big.Int)(x).Cmp((*big.Int)(y)) <= 0
}

func (x *Int) Lt(y *Int) bool {
	return (*big.Int)(x).Cmp((*big.Int)(y)) < 0
}

func (x *Int) Eq(y *Int) bool {
	return (*big.Int)(x).Cmp((*big.Int)(y)) == 0
}

func (x *Int) Add(y *Int) *Int {
	if x == nil {
		x = NewInt(0)
	}

	if y == nil {
		y = NewInt(0)
	}

	return (*Int)(big.NewInt(0).Add((*big.Int)(x), (*big.Int)(y)))
}

func (x *Int) Sub(y *Int) *Int {
	if x == nil {
		x = NewInt(0)
	}

	if y == nil {
		y = NewInt(0)
	}

	return (*Int)(big.NewInt(0).Sub((*big.Int)(x), (*big.Int)(y)))
}

func (x *Int) Neg() *Int {
	return (*Int)(big.NewInt(0).Neg((*big.Int)(x)))
}

func (x *Int) String() string {
	return (*big.Int)(x).String()
}

func (x *Int) MarshalJSON() ([]byte, error) {
	return (*big.Int)(x).MarshalJSON()
}

func (x *Int) UnmarshalJSON(b []byte) error {
	return (*big.Int)(x).UnmarshalJSON(b)
}

func (x *Int) MarshalText() ([]byte, error) {
	return (*big.Int)(x).MarshalText()
}

func (x *Int) Int64() int64 {
	return (*big.Int)(x).Int64()
}
