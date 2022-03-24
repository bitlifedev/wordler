package dictionary

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Words struct {
	words []word
	hash  []int
}

type word struct {
	word   string
	hash   []int
	active bool
}

func open(file string) (*os.File, error) {
	// Open our jsonFile
	contents, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer func(contents *os.File) {
		if err != nil {
			log.Fatal(err)
		}
	}(contents)
	fmt.Println("Successfully Opened dictionary")
	return contents, nil
}

func Load(dictionary string) (Words, error) {

	// read our opened xmlFile as a byte array.
	var WS Words
	file, _ := open(dictionary)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		var W word
		W.word = scanner.Text()
		W.hash = hashString(W.word)
		W.active = true
		WS.words = append(WS.words, W)
		for i := range W.hash {
			WS.hash[i] = WS.hash[i] + W.hash[i]
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error occured scanning: %v", err)
	}
	return WS, nil
}

func hashString(S string) []int {

	empty := [26]int{}
	myIntArray := empty[0:25]

	for i := range S {
		s := strings.ToUpper(string(S[i]))
		index := s[0] - 65
		if index < 25 {
			myIntArray[index]++
		}

	}

	return myIntArray
}
