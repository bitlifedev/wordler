package dictionary

import (
	"io/fs"
	"testing"
)

type expected struct {
	file  *fs.File
	error error
}

func TestOpen(t *testing.T) {
	var tests = []struct {
		input string
	}{
		{"../assets/test.dic"},
	}
	for _, test := range tests {
		if output, _ := open(test.input); output == nil {
			t.Errorf("Test Failed: %v inputted,received: %v", test.input, output)
		}
	}

}

func TestTableReverse(t *testing.T) {

}
