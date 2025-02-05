package wordcounter

import (
	"testing"
	"reflect"
)


func initializeEmptyCounter() *WordCounter{
	return NewWordCounter()
}

func initializeCounter() *WordCounter{
	c := NewWordCounter()
	c.Add("livro")
	c.Add("livro")
	c.Add("cinema")
	c.Add("praia")
	c.Add("livro")
	c.Add("cinema")
	return c
}


func TestNewWordCounter(t *testing.T) {
	wordCounter := NewWordCounter()

	qtdWords := len(wordCounter.countMap)

	if qtdWords != 0 {
		t.Fail()
	}
}

func TestAdd(t *testing.T){
	counter := initializeEmptyCounter()

	counter.Add("livro")
	counter.Add("livro")
	counter.Add("cinema")

	t.Log(counter.countMap)
	if counter.countMap["livro"] != 2 {
		t.Fail()
	}

	if counter.countMap["cinema"] != 1 {
		t.Fail()
	}
	
}

func TestMaxKeyValue(t *testing.T){
	counter := initializeEmptyCounter()
	counter.Add("livro")
	counter.Add("livro")
	counter.Add("cinema")	

	if counter.maxKey != "livro" || counter.maxValue != 2 {
		t.Fail()
	}
}

func TestHowManyWords(t *testing.T) {
	counter := initializeEmptyCounter()
	counter.Add("livro")
	counter.Add("livro")
	counter.Add("cinema")	

	if counter.howManyWords != 3 {
		t.Fail()
	}
	
}

func TestSortedKeys(t *testing.T) {
	counter := initializeCounter()
	var expectedKeys, keys []string
	keys = counter.GetKeysByOrderDesc()
	expectedKeys = []string{"livro", "cinema", "praia"}

	t.Log(keys)

	if !reflect.DeepEqual(keys, expectedKeys) {
		t.Fail()
	}
}

func TestCountWords(t *testing.T) {
	counter := initializeEmptyCounter()
	words := []string{"livro", "livro", "cinema", "praia", "livro", "cinema"}

	counter.Count(&words)

	expectedCount := map[string]int{
		"livro": 3,
		"cinema": 2,
		"praia": 1,
	}

	if !reflect.DeepEqual(expectedCount, counter.countMap) {
		t.Fail()
	}
}

func TestCleanText(t *testing.T) {
	text := ",o rato roeu.?a!..._roupa-:;\ndo rei\r\t\ufeffde--«roma»"
	cleanedText := CleanText(text)
	// t.Log(cleanedText)
	expected := " o rato roeu  a     roupa    do rei   de   roma "
	if cleanedText != expected {
		t.Fail()
	}
}

func TestSplitText(t *testing.T) {
	text := "o rato roeu a     roupa do rei de roma"
	splittedText := SplitText(text)
	expected := []string{"o", "rato", "roeu", "a", "roupa", "do", "rei", "de", "roma"}
	if !reflect.DeepEqual(expected, splittedText) {
		t.Fail()
	}
}
