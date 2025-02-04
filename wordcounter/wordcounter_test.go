package wordcounter

import "testing"

func initializeEmptyCounter() *WordCounter{
	return NewWordCounter()
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