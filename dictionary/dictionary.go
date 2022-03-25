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
	Map   map[string]Stats
}

type Word struct {
	Value string
	Map   map[string]Stats
}
type Stats struct {
	Count int
	P     float64
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
	dictionary.Map = make(map[string]Stats)
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
			dictionary.Map[i] = Stats{dictionary.Map[i].Count + w.Map[i].Count, -1}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error occured scanning: %v", err)
	}
	return dictionary, nil
}

func mapString(S string) map[string]Stats {
	m := make(map[string]Stats)
	for i := range S {
		s := strings.ToUpper(string(S[i]))
		m[s] = Stats{m[s].Count + 1, (float64(m[s].Count + 1)) / float64(len(S))}
	}
	return m
}
