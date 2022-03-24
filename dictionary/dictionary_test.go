package dictionary

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOpen(t *testing.T) {

	var tests = []struct {
		input    string
		expected *os.File
		err      string
	}{
		{"../assets/fail.dic", nil, "open ../assets/fail.dic: no such file or directory"},
		{"../assets/test.dic", nil, ""},
	}

	for _, test := range tests {
		_, err := open(test.input)
		if err != nil {
			assert.Error(t, err)
			{
				assert.Equal(t, test.err, err.Error())
			}
		}
	}
}

func TestHashString(t *testing.T) {
	var tests = []struct {
		input    string
		expected []int
	}{
		{"A", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"a", []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"Aa", []int{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"Ab", []int{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{" ", []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"", []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{"1", []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	}
	for _, test := range tests {
		output := hashString(test.input)
		for p := 0; p < len(output); p++ {
			if output[p] != test.expected[p] {
				t.Errorf("Test Failed: input: %v expected: %v output:%v", test.input, test.expected, output)
			}
			fmt.Printf("%v", output[p])
		}
		fmt.Println(" √")
	}
}
