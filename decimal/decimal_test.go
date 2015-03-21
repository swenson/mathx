package decimal

import "testing"

func TestInputOutput(t *testing.T) {
	cases := []string{"0.0", "-1", "-1.0", "-1.01", "-123.456", "123123123123444444.412341923480192384901"}
	for _, c := range cases {
		d, err := New(c)
		if err != nil {
			t.Fatalf("Error reading in perfectly ordinary decimal %s: %s", c, err.Error())
		}
		out := d.String()
		if out != c {
			t.Errorf("Decimal string of %s was unexpectedly %s", c, out)
		}
	}
}
