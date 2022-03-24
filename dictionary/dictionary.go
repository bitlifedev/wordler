package dictionary

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Dictionary struct {
	Words []Word
	Map   map[string]int
}

type Word struct {
	Value string
	Map   map[string]int
}

func open(file string) (*os.File, error) {
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

func Load(file string) (Dictionary, error) {

	var dictionary Dictionary
	dictionary.Map = make(map[string]int)
	content, _ := open(file)
	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		var w Word
		var s = scanner.Text()
		w.Value = s
		w.Map = mapString(s)
		dictionary.Words = append(dictionary.Words, w)
		for i := range w.Map {
			dictionary.Map[i] = dictionary.Map[i] + w.Map[i]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error occured scanning: %v", err)
	}
	return dictionary, nil
}

func mapString(S string) map[string]int {
	m := make(map[string]int)
	for i := range S {
		s := strings.ToUpper(string(S[i]))
		m[s] = m[s] + 1
	}
	return m
}
