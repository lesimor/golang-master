package main

import (
	"fmt"
	"testing"
)

type testConfig struct {
	args []string
	err  bool
	config
}

func TestParseArgs(t *testing.T) {
	tests := []testConfig{
		{
			args:   []string{"-h"},
			err:    false,
			config: config{printUsage: true, numTimes: 0},
		},
		{
			args:   []string{"10"},
			err:    false,
			config: config{printUsage: false, numTimes: 10},
		},
		{
			args:   []string{"abc"},
			err:    true,
			config: config{printUsage: false, numTimes: 0},
		},
		{
			args:   []string{"1", "foo"},
			err:    true,
			config: config{printUsage: false, numTimes: 0},
		},
	}

	for _, tc := range tests {
		c, err := parseArgs(tc.args)
		if tc.err == false && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
		if tc.err == true && err == nil {
			t.Errorf("Expected error exists but no exception occured")
			fmt.Println(tc)
		}

		if c.numTimes != tc.numTimes {
			t.Errorf("Expected numTimes to be: %v, got: %v\n", tc.numTimes, c.numTimes)
		}
	}
}
