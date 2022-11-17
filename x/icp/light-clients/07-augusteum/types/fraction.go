package types

import (
	tmmath "github.com/creatachain/augusteum/libs/math"
	"github.com/creatachain/augusteum/light"
)

// DefaultTrustLevel is the augusteum light client default trust level
var DefaultTrustLevel = NewFractionFromTm(light.DefaultTrustLevel)

// NewFractionFromTm returns a new Fraction instance from a tmmath.Fraction
func NewFractionFromTm(f tmmath.Fraction) Fraction {
	return Fraction{
		Numerator:   f.Numerator,
		Denominator: f.Denominator,
	}
}

// ToAugusteum converts Fraction to tmmath.Fraction
func (f Fraction) ToAugusteum() tmmath.Fraction {
	return tmmath.Fraction{
		Numerator:   f.Numerator,
		Denominator: f.Denominator,
	}
}
