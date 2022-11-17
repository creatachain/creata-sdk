package types

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	cta  = "cta"  // 1 (base denom unit)
	mcta = "mcta" // 10^-3 (milli)
	ucta = "ucta" // 10^-6 (micro)
	ncta = "ncta" // 10^-9 (nano)
)

type internalDenomTestSuite struct {
	suite.Suite
}

func TestInternalDenomTestSuite(t *testing.T) {
	suite.Run(t, new(internalDenomTestSuite))
}

func (s *internalDenomTestSuite) TestRegisterDenom() {
	ctaUnit := OneDec() // 1 (base denom unit)

	s.Require().NoError(RegisterDenom(cta, ctaUnit))
	s.Require().Error(RegisterDenom(cta, ctaUnit))

	res, ok := GetDenomUnit(cta)
	s.Require().True(ok)
	s.Require().Equal(ctaUnit, res)

	res, ok = GetDenomUnit(mcta)
	s.Require().False(ok)
	s.Require().Equal(ZeroDec(), res)

	// reset registration
	baseDenom = ""
	denomUnits = map[string]Dec{}
}

func (s *internalDenomTestSuite) TestConvertCoins() {
	ctaUnit := OneDec() // 1 (base denom unit)
	s.Require().NoError(RegisterDenom(cta, ctaUnit))

	mctaUnit := NewDecWithPrec(1, 3) // 10^-3 (milli)
	s.Require().NoError(RegisterDenom(mcta, mctaUnit))

	uctaUnit := NewDecWithPrec(1, 6) // 10^-6 (micro)
	s.Require().NoError(RegisterDenom(ucta, uctaUnit))

	nctaUnit := NewDecWithPrec(1, 9) // 10^-9 (nano)
	s.Require().NoError(RegisterDenom(ncta, nctaUnit))

	res, err := GetBaseDenom()
	s.Require().NoError(err)
	s.Require().Equal(res, ncta)
	s.Require().Equal(NormalizeCoin(NewCoin(ucta, NewInt(1))), NewCoin(ncta, NewInt(1000)))
	s.Require().Equal(NormalizeCoin(NewCoin(mcta, NewInt(1))), NewCoin(ncta, NewInt(1000000)))
	s.Require().Equal(NormalizeCoin(NewCoin(cta, NewInt(1))), NewCoin(ncta, NewInt(1000000000)))

	coins, err := ParseCoinsNormalized("1cta,1mcta,1ucta")
	s.Require().NoError(err)
	s.Require().Equal(coins, Coins{
		Coin{ncta, NewInt(1000000000)},
		Coin{ncta, NewInt(1000000)},
		Coin{ncta, NewInt(1000)},
	})

	testCases := []struct {
		input  Coin
		denom  string
		result Coin
		expErr bool
	}{
		{NewCoin("foo", ZeroInt()), cta, Coin{}, true},
		{NewCoin(cta, ZeroInt()), "foo", Coin{}, true},
		{NewCoin(cta, ZeroInt()), "FOO", Coin{}, true},

		{NewCoin(cta, NewInt(5)), mcta, NewCoin(mcta, NewInt(5000)), false},       // cta => mcta
		{NewCoin(cta, NewInt(5)), ucta, NewCoin(ucta, NewInt(5000000)), false},    // cta => ucta
		{NewCoin(cta, NewInt(5)), ncta, NewCoin(ncta, NewInt(5000000000)), false}, // cta => ncta

		{NewCoin(ucta, NewInt(5000000)), mcta, NewCoin(mcta, NewInt(5000)), false},       // ucta => mcta
		{NewCoin(ucta, NewInt(5000000)), ncta, NewCoin(ncta, NewInt(5000000000)), false}, // ucta => ncta
		{NewCoin(ucta, NewInt(5000000)), cta, NewCoin(cta, NewInt(5)), false},            // ucta => cta

		{NewCoin(mcta, NewInt(5000)), ncta, NewCoin(ncta, NewInt(5000000000)), false}, // mcta => ncta
		{NewCoin(mcta, NewInt(5000)), ucta, NewCoin(ucta, NewInt(5000000)), false},    // mcta => ucta
	}

	for i, tc := range testCases {
		res, err := ConvertCoin(tc.input, tc.denom)
		s.Require().Equal(
			tc.expErr, err != nil,
			"unexpected error; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
		s.Require().Equal(
			tc.result, res,
			"invalid result; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
	}

	// reset registration
	baseDenom = ""
	denomUnits = map[string]Dec{}
}

func (s *internalDenomTestSuite) TestConvertDecCoins() {
	ctaUnit := OneDec() // 1 (base denom unit)
	s.Require().NoError(RegisterDenom(cta, ctaUnit))

	mctaUnit := NewDecWithPrec(1, 3) // 10^-3 (milli)
	s.Require().NoError(RegisterDenom(mcta, mctaUnit))

	uctaUnit := NewDecWithPrec(1, 6) // 10^-6 (micro)
	s.Require().NoError(RegisterDenom(ucta, uctaUnit))

	nctaUnit := NewDecWithPrec(1, 9) // 10^-9 (nano)
	s.Require().NoError(RegisterDenom(ncta, nctaUnit))

	res, err := GetBaseDenom()
	s.Require().NoError(err)
	s.Require().Equal(res, ncta)
	s.Require().Equal(NormalizeDecCoin(NewDecCoin(ucta, NewInt(1))), NewDecCoin(ncta, NewInt(1000)))
	s.Require().Equal(NormalizeDecCoin(NewDecCoin(mcta, NewInt(1))), NewDecCoin(ncta, NewInt(1000000)))
	s.Require().Equal(NormalizeDecCoin(NewDecCoin(cta, NewInt(1))), NewDecCoin(ncta, NewInt(1000000000)))

	coins, err := ParseCoinsNormalized("0.1cta,0.1mcta,0.1ucta")
	s.Require().NoError(err)
	s.Require().Equal(coins, Coins{
		Coin{ncta, NewInt(100000000)},
		Coin{ncta, NewInt(100000)},
		Coin{ncta, NewInt(100)},
	})

	testCases := []struct {
		input  DecCoin
		denom  string
		result DecCoin
		expErr bool
	}{
		{NewDecCoin("foo", ZeroInt()), cta, DecCoin{}, true},
		{NewDecCoin(cta, ZeroInt()), "foo", DecCoin{}, true},
		{NewDecCoin(cta, ZeroInt()), "FOO", DecCoin{}, true},

		// 0.5cta
		{NewDecCoinFromDec(cta, NewDecWithPrec(5, 1)), mcta, NewDecCoin(mcta, NewInt(500)), false},       // cta => mcta
		{NewDecCoinFromDec(cta, NewDecWithPrec(5, 1)), ucta, NewDecCoin(ucta, NewInt(500000)), false},    // cta => ucta
		{NewDecCoinFromDec(cta, NewDecWithPrec(5, 1)), ncta, NewDecCoin(ncta, NewInt(500000000)), false}, // cta => ncta

		{NewDecCoin(ucta, NewInt(5000000)), mcta, NewDecCoin(mcta, NewInt(5000)), false},       // ucta => mcta
		{NewDecCoin(ucta, NewInt(5000000)), ncta, NewDecCoin(ncta, NewInt(5000000000)), false}, // ucta => ncta
		{NewDecCoin(ucta, NewInt(5000000)), cta, NewDecCoin(cta, NewInt(5)), false},            // ucta => cta

		{NewDecCoin(mcta, NewInt(5000)), ncta, NewDecCoin(ncta, NewInt(5000000000)), false}, // mcta => ncta
		{NewDecCoin(mcta, NewInt(5000)), ucta, NewDecCoin(ucta, NewInt(5000000)), false},    // mcta => ucta
	}

	for i, tc := range testCases {
		res, err := ConvertDecCoin(tc.input, tc.denom)
		s.Require().Equal(
			tc.expErr, err != nil,
			"unexpected error; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
		s.Require().Equal(
			tc.result, res,
			"invalid result; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
	}

	// reset registration
	baseDenom = ""
	denomUnits = map[string]Dec{}
}

func (s *internalDenomTestSuite) TestDecOperationOrder() {
	dec, err := NewDecFromStr("11")
	s.Require().NoError(err)
	s.Require().NoError(RegisterDenom("unit1", dec))
	dec, err = NewDecFromStr("100000011")
	s.Require().NoError(RegisterDenom("unit2", dec))

	coin, err := ConvertCoin(NewCoin("unit1", NewInt(100000011)), "unit2")
	s.Require().NoError(err)
	s.Require().Equal(coin, NewCoin("unit2", NewInt(11)))

	// reset registration
	baseDenom = ""
	denomUnits = map[string]Dec{}
}
