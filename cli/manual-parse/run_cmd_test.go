package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		c      config
		input  string
		output string
		err    bool
	}{
		{
			c:      config{printUsage: true},
			output: usageString,
		},
		{
			c:      config{numTimes: 10},
			input:  "",
			output: "Your name please? Press the Enter key when done.\n",
			err:    true,
		},
		{
			c:      config{numTimes: 10},
			input:  "kang",
			output: "Your name please? Press the Enter key when done.\n" + strings.Repeat("nice to meet you kang\n", 10),
			err:    false,
		},
	}

	byteBuf := new(bytes.Buffer)
	for _, tc := range tests {
		rd := strings.NewReader(tc.input)
		err := runCmd(rd, byteBuf, tc.c)
		if err != nil && tc.err == false {
			t.Errorf("Expected nil error, got: %v\n", err)
		}
		if err == nil && tc.err == true {
			t.Errorf("Expected error exists but no exception occured")
		}
		if byteBuf.String() != tc.output {
			t.Errorf("Expected output: %v, got: %v\n", tc.output, byteBuf.String())
		}
		byteBuf.Reset()
	}
}
