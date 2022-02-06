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

func TestTimesThree(t *testing.T) {
	tests := []struct {
		name     string
		send     int
		channel  chan int
		response int
	}{
		{
			"Test 2 *3",
			2,
			make(chan int),
			6,
		},
		{
			"Test 3 * 3",
			3,
			make(chan int),
			9,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			// Setup
			var result int
			// Execute
			go timesThree(testcase.send, testcase.channel)
			// Check
			result = <-testcase.channel
			if result != testcase.response {
				t.Errorf("Failure to calculate correct response, expected: %d -- response: %d", testcase.response, result)
			}
		})

	}
}
