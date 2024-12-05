package fn

import "testing"

func TestRoundDownToTwoDecimalPlaces(t *testing.T) {
	num := 0.000369875321
	t.Logf("0-> %+v", RoundDownToTwoDecimalPlaces(num, 0))
	t.Logf("2-> %+v", RoundDownToTwoDecimalPlaces(num))
	t.Logf("4-> %+v", RoundDownToTwoDecimalPlaces(num, 4))
	t.Logf("6-> %+v", RoundDownToTwoDecimalPlaces(num, 6))
	t.Logf("20-> %+v", RoundDownToTwoDecimalPlaces(num, 20))
}
