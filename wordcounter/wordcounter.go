package wordcounter

type WordCounter struct {
	countMap map[string]int // Must have a mapping from string
}

func NewWordCounter() *WordCounter {
	return &WordCounter{};
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
}

