package main

import (
	"fmt"
	"sort"
)

type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }
func (r runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func sortString(s string) string {
	r := runes(s)
	sort.Sort(r)
	return string(r)
}

func areAnagrams(s, t string) bool {
	return sortString(s) == sortString(t)
}

func main() {
	m := map[string]string{
		"listen": "silent",
		"armpit": "impart",
		"cat":    "dog",
		"lazy":   "bear",
	}

	for k, v := range m {
		if areAnagrams(k, v) {
			fmt.Printf("%s and %s are anagrams! :)\n", k, v)
		} else {
			fmt.Printf("%s and %s are not anagrams! :(\n", k, v)
		}
	}
}
