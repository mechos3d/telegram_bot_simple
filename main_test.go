package main

import (
	"testing"
)

func TestReplyMessageText(t *testing.T) {

	in := "/start"
	want := "you typed /start"
	got := replyMessageText(in)

	if got != want {
		t.Fatalf(`FAIL: replyMessageText(%s) = %s want %s`, in, got, want)
	}
	t.Log("PASS:", "replyMessageText")
}
