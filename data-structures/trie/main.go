package main

// The number of children nodes each trie Node can contain
const (
	NumNodes = 26
)

type Node struct {
	children [NumNodes]*Node
	isEnd bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	return &Trie{
		root: &Node{
			isEnd: false,
		},
	}
}

func (t *Trie) Insert(word string) {
	wordLength := len(word)
	node := t.root
	// Traverse the trie comparing each letter
	for i := 0; i < wordLength; i++ {
		// Using the character index ensures the same index is used for each character in each array
		// This is kinda like hashing a bit actually I think, want to look into that a little more
		charIndex := word[i] -'a' // this gets the character as an index starting a at index 0
		if node.children[charIndex] == nil { // insert new node
			node.children[charIndex] = &Node{}
		}
		node = node.children[charIndex]
	}
	// set final node to be the end of a word
	node.isEnd = true;
}

func (t *Trie) Search(word string) bool {
	wordLength := len(word)
	node := t.root
	// Traverse the trie comparing each letter
	for i := 0; i < wordLength; i++ {
		// Using the character index ensures the same index is used for each character in each array
		// This is kinda like hashing a bit actually I think, want to look into that a little more
		charIndex := word[i] -'a' // this gets the character as an index starting a at index 0
		if node.children[charIndex] == nil {
			return false
		}
		node = node.children[charIndex]
	}
	return node.isEnd
}


func main() {}
