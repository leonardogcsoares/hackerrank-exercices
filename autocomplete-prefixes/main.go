package main

import "fmt"

/*
The problem is given an array of words:
[abc, aba, bca, bba]

You type a prefix:
prefix = a
and it should return the array of words that start with the prefix:
output = [abc, aba]

One solution would be to iterate through the array and check if:
prefix == word[:len(prefix)]

But this would be done in O(n*k) time since we would need to iterate through the array completely checking each individual
element. WHere K is the len of the substring.

Using tries we could build a tree of all the possiblities of substrings and words.

*/

type node struct {
	Val      string
	IsWord   bool
	Children map[string]*node
}

var trie *node

func main() {
	trie = &node{
		Val:      "",
		Children: make(map[string]*node),
	}

	words := []string{"aba", "abc", "bca", "bba", "acb", "abb", "acc"}
	for _, word := range words {
		insertWord(word)
	}

	fmt.Println(getWordsForPrefix("ab"))
}

func insertWord(word string) {
	curr := trie
	for k, c := range word {
		cs := string(c)

		nextNode, ok := curr.Children[cs]
		if !ok {
			nextNode = &node{
				Val:      word[:k+1],
				Children: make(map[string]*node),
			}
			curr.Children[cs] = nextNode
		}

		curr = nextNode
		if k == len(word)-1 {
			curr.IsWord = true
		}
	}
}

func getWordsForPrefix(prefix string) []string {
	var words []string

	curr := trie

	for _, c := range prefix {
		cs := string(c)
		nextNode, ok := curr.Children[cs]
		if ok {
			curr = nextNode
		} else {
			return words
		}
	}

	findAllChildWords(curr, &words)

	return words
}

func findAllChildWords(curr *node, words *[]string) {
	if curr.IsWord {
		*words = append(*words, curr.Val)
	}

	for _, v := range curr.Children {
		findAllChildWords(v, words)
	}
}
