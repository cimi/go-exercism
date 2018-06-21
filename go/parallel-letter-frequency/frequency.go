package letter

import "sync"

type FreqMap map[rune]int

func (f FreqMap) Add(other FreqMap) {
	for k, v := range other {
		f[k] += v
	}
}

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(input []string) FreqMap {
	m := FreqMap{}

	c := make(chan FreqMap, len(input))
	var wg sync.WaitGroup
	for _, s := range input {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			c <- Frequency(s)
		}(s)
	}
	wg.Wait()
	close(c)

	for r := range c {
		m.Add(r)
	}

	return m
}
