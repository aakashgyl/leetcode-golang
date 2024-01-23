package main

type Trie struct {
	data      map[byte]*Trie
	endOfWord bool
}

func Constructor() Trie {
	return Trie{
		data: make(map[byte]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	curr := this
	for i := 0; i < len(word); i++ {
		var next *Trie
		if next = curr.data[word[i]]; next == nil {
			next = &Trie{
				data: make(map[byte]*Trie),
			}
			curr.data[word[i]] = next
		}
		curr = next
	}
	curr.endOfWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for i := 0; i < len(word); i++ {
		var next *Trie
		if next = curr.data[word[i]]; next == nil {
			return false
		}
		curr = next
	}
	return curr.endOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for i := 0; i < len(prefix); i++ {
		var next *Trie
		if next = curr.data[prefix[i]]; next == nil {
			return false
		}
		curr = next
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
