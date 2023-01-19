package main

import (
	"testing"
)

func TestValidateArgs(t *testing.T) {
	for _, tc := range []struct {
		config
		err bool
	}{
		{
			config: config{},
			err:    true,
		},
		{
			config: config{numTimes: -1},
			err:    true,
		},
		{
			config: config{numTimes: 10},
			err:    false,
		},
	} {
		err := validateArgs(tc.config)
		if tc.err == false && err != nil {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
		if tc.err == true && err == nil {
			t.Errorf("Expected error exists but no exception occured")
		}
	}
}
