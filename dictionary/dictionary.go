package dictionary

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"wordler/words"
)

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

func Load(file string) (*words.WordlerDictionary, error) {

	content, _ := open(file)
	scanner := bufio.NewScanner(content)
	scanner.Split(bufio.ScanWords)
	var dict = new(words.WordlerDictionary)
	dict.Map = map[string]words.Stats{}
	for scanner.Scan() {
		w := new(words.Word)
		w.Value = scanner.Text()
		w.Map = mapString(w.Value)
		updateWordlerDictionary(dict, w)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error occured scanning: %v", err)
		return nil, err
	}
	return dict, nil
}

func updateWordlerDictionary(dict *words.WordlerDictionary, w *words.Word) {
	dict.Words = append(dict.Words, *w)
	var stats words.Stats
	stats.Count = 0
	for s := range w.Map {
		stats.Count = dict.Map[s].Count + w.Map[s].Count
		dict.Map[s] = stats
	}
}

func mapString(S string) map[string]words.Stats {
	//m := make(map[string]Stats)
	var m = map[string]words.Stats{}
	var stats words.Stats
	stats.Count = 0
	for i := range S {
		s := strings.ToUpper(string(S[i]))
		stats.Count = m[s].Count + 1
		m[s] = stats
	}
	return m
}
