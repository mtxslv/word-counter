package wordcounter

import "testing"

func TestNewWordCounter(t *testing.T) {
	wordCounter := NewWordCounter()

	qtdWords := len(wordCounter.countMap)

	if qtdWords != 0 {
		t.Fail()
	}
}