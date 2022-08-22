package designAddAndSearchWordsDataStructure

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	node := this
	for _, c := range word {
		if node.children[int(c-'a')] == nil {
			node.children[int(c-'a')] = &WordDictionary{}
		}
		node = node.children[int(c-'a')]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	node := this
	var res bool
	for i, c := range word {
		if c == '.' {
			for j := 0; j < 26; j++ {
				if node.children[j] != nil {
					res = node.children[j].Search(word[i+1:]) || res
				}
			}
			return res
		}
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

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
