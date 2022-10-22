package main

import "testing"

func TestMain(m *testing.T) {
	want := "hello"

	if got := hello(); got != want {
		m.Fatalf(`hello("") = %q, %v, want "", error`, got, want)
	}
}
