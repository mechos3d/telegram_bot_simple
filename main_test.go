package main

import (
	"testing"
)

func TestProcessMessage(t *testing.T) {

	in := "/start"
	want := "you typed /start"
	got := processMessage(in)

	if got != want {
		t.Fatalf(`FAIL: processMessage(%s) = %s want %s`, in, got, want)
	}
	t.Log("PASS:", "processMessage")
}
