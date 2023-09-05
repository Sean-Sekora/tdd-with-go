package wordz

type WordRepository interface {
	FetchWordByNumber(int) string
	HighestWordNumber() int
}
