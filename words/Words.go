package words

type WordlerDictionary struct {
	Words []Word
	Map   map[string]Stats
}

type Word struct {
	Value string
	Map   map[string]Stats
}
type Stats struct {
	Count int
}
