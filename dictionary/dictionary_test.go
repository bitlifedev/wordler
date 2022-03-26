package dictionary

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"wordler/words"
)

func TestLoad(t *testing.T) {

	var tests = []struct {
		input    string
		expected words.WordlerDictionary
		err      string
	}{
		{"../assets/test.dic", words.WordlerDictionary{}, ""},
	}
	for _, test := range tests {
		received, err := Load(test.input)
		if err != nil {
		}
		if &received == nil {
			t.Errorf("Empty Dictionary")
		}
	}
}

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
		expected int
	}{
		{input: "A", expected: 2},
		{input: "Aa", expected: 2},
		{input: "Ab", expected: 2},
		{input: " ", expected: 2},
		{input: "1", expected: 2},
		{input: "", expected: 2},
	}
	for _, test := range tests {
		output := mapString(test.input)
		fmt.Printf("%v √ \n", output)
	}
}
