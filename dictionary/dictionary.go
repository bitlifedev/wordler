package dictionary

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Words struct {
	Words []Word
	Hash  [26]int
}

type Word struct {
	Word string
	Hash []int
}

func open(file string) (*os.File, error) {
	// Open our jsonFile
	contents, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func(contents *os.File) {
		if err != nil {
			fmt.Println(err)
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
		var W Word
		W.Word = scanner.Text()
		W.Hash = hashString(W.Word)
		WS.Words = append(WS.Words, W)
		for i := range W.Hash {
			WS.Hash[i] = WS.Hash[i] + W.Hash[i]
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
