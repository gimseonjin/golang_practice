package main1

import "testing"

func TestMain(m *testing.T) {
	want := "μλ!, μΈμ."

	if got := hello(); got != want {
		m.Fatalf(`hello("") = %q, %v, want "", error`, got, want)
	}
}
