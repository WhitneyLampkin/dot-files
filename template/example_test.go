package main

import "testing"

func TestHelloWorld(t *testing.T) {
        // t.Fatal("not implemented")
}

func TestExample(t *testing.T) {
        Run()
        got := ""
        expect := ""
        if got != expect {
                t.Errorf("got %q want %q", got, expect)
        }
}
