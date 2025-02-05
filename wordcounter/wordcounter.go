package wordcounter

import (
	"slices"
)

type WordCounter struct {
	countMap map[string]int // Must have a mapping from string
	maxValue int 
	maxKey string
	howManyWords int
}

func NewWordCounter() *WordCounter {
	return &WordCounter{
		countMap: make(map[string]int),
		maxValue: -1,
		maxKey: "",
		howManyWords: 0,
	};
}

func (c *WordCounter) Add (k string) {
	// Add a certain k string to the counting.
	// If k already exists, increase the counting.
	// Otherwise, count as 1.

	// Is there a key k on the counting?
	if _, ok := c.countMap[k]; ok {
		c.countMap[k]++; // if so, add 1
	} else {
		c.countMap[k] = 1; // not the case, add
	}

	// keep the word that appears the most
	if c.countMap[k] > c.maxValue {
		c.maxValue = c.countMap[k] 
		c.maxKey = k
	}

	// increase word count
	c.howManyWords++
}

func (c *WordCounter) getKeysByOrderDesc() []string {

	// Put all keys in an array
	allKeys := make([]string,len(c.countMap)) // empty
	var i int = 0
	for k := range c.countMap {
		allKeys[i] = k
		i++
	}

	// anonymous func for comparison
	keyCmp := func(k1, k2 string) int {
		return c.countMap[k2] - c.countMap[k1]
	}
	// Sort keys desc
	slices.SortFunc(
		allKeys,
		keyCmp,	
	)

	return allKeys
}