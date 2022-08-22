package implementTriePrefixTree

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, c := range word {
		if node.children[int(c-'a')] == nil {
			node.children[int(c-'a')] = &Trie{}
		}
		node = node.children[int(c-'a')]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		if node.children[int(c-'a')] == nil {
			return false
		}
		node = node.children[int(c-'a')]
	}
	if node.isEnd {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		if node.children[int(c-'a')] == nil {
			return false
		}
		node = node.children[int(c-'a')]
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
