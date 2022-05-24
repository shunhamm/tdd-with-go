package main

import "testing"

func TestHellow(t *testing.T) {
	got := Hello()
	want := "Hello world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
