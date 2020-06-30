// 3.12 - Write a function that reports whether two strings are anagrams of each
// other, that is, they contain the same letters in a different order.
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

func areAnagramsTwo(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[byte]int)
	for i, _ := range s {
		m[s[i]]++
		m[t[i]]--
	}

	for k := range m {
		if m[k] != 0 {
			return false
		}
	}
	return true
}

// This assumes lowercase ASCII characters only
func areAnagramsThree(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var sCount, tCount [26]int

	for i := 0; i < len(s); i++ {
		sCount[s[i]-'a']++
		tCount[t[i]-'a']++
	}

	return sCount == tCount
}

func main() {
	m := map[string]string{
		"listen": "silent",
		"armpit": "impart",
		"cat":    "dog",
		"lazy":   "bear",
	}

	fmt.Println("Using string sorting..")
	for k, v := range m {
		fmt.Println(k, v, areAnagrams(k, v))
	}
	fmt.Println("")
	fmt.Println("Using byte frequency map..")
	for k, v := range m {
		fmt.Println(k, v, areAnagramsTwo(k, v))
	}
	fmt.Println("")
	fmt.Println("Using char count array..")
	for k, v := range m {
		fmt.Println(k, v, areAnagramsThree(k, v))
	}
}
