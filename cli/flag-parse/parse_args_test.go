package main

import (
	"bytes"
	"fmt"
	"testing"
)

type testConfig struct {
	args     []string
	err      bool
	numTimes int
}

func TestParseArgs(t *testing.T) {
	tests := []testConfig{
		{
			args:     []string{"-h"},
			err:      true,
			numTimes: 0,
		},
		{
			args:     []string{"-n", "10"},
			err:      false,
			numTimes: 10,
		},
		{
			args:     []string{"abc"},
			err:      true,
			numTimes: 0,
		},
		{
			args:     []string{"1", "foo"},
			err:      true,
			numTimes: 0,
		},
	}
	bytesBuf := new(bytes.Buffer)
	for _, tc := range tests {
		c, err := parseArgs(bytesBuf, tc.args)
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
		bytesBuf.Reset()
	}
}
