package main

import "testing"

func TestMain(m *testing.T) {
	want := "안녕!, 세상."

	if got := hello(); got != want {
		m.Fatalf(`hello("") = %q, %v, want "", error`, got, want)
	}
}
