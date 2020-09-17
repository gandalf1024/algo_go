package cycle

import "testing"

func Test_Add(t *testing.T) {
	l := New()
	l.Add(12, "BBB")
	l.Add(14, "CCC")
	l.Add(17, "DDD")
	l.Add(88, "FFF")
	l.Add(56, "EEE")
	l.Add(2, "AAA")

	l.List()
}
