package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	Run()
	got := ""
	expect := ""
	if got != expect {
		t.Errorf("got %q want %q", got, expect)
	}
}
