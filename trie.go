package trie

type Trie struct {
	root *node
	size int
}
type node struct {
	r        rune
	path     string
	term     bool
	meta     interface{}
	parent   *node
	children map[rune]*node
}

func New() *Trie {
	return &Trie{
		root: &node{
			r:        0,
			term:     false,
			meta:     nil,
			children: make(map[rune]*node),
		},
		size: 0,
	}
}

func (t *Trie) Size() int {
	return t.size
}

func fuzzyCollect(n *node, fuzzy []rune, index int) []string {
	var s []string

	if n.term {
		if index == len(fuzzy) {
			s = append(s, n.path)
		}
	}

	for _, v := range n.children {
		var inc int
		if index != len(fuzzy) && v.r == fuzzy[index] {
			inc = 1
		}
		s = append(s, fuzzyCollect(v, fuzzy, index+inc)...)
	}
	return s
}
func (t *Trie) FuzzySearch(fuzzy string) []string {
	return fuzzyCollect(t.root, []rune(fuzzy), 0)
}

func fuzzyCollectWithElem(n *node, fuzzy []rune, index int) []strKV {
	var s []strKV

	if n.term {
		if index == len(fuzzy) {
			s = append(s, strKV{
				Key:   n.path,
				Value: n.meta,
			})
		}
	}

	for _, v := range n.children {
		var inc int
		if index != len(fuzzy) && v.r == fuzzy[index] {
			inc = 1
		}
		s = append(s, fuzzyCollectWithElem(v, fuzzy, index+inc)...)
	}
	return s
}
func (t *Trie) FuzzySearchWithElem(fuzzy string) []strKV {
	return fuzzyCollectWithElem(t.root, []rune(fuzzy), 0)
}

func collect(n *node) []string {
	var s []string
	if n.term {
		s = append(s, n.path)
	}

	for _, v := range n.children {
		s = append(s, collect(v)...)
	}
	return s
}
func (t *Trie) PrefixSearch(pre string) []string {
	nByQuery := findNode(t.root, []rune(pre))
	if nByQuery == nil {
		return nil
	}

	return collect(nByQuery)
}

func collectWithElem(node *node) []strKV {
	var s []strKV
	if node.term {
		s = append(s, strKV{
			Key:   node.path,
			Value: node.meta,
		})
	}

	for _, v := range node.children {
		s = append(s, collectWithElem(v)...)
	}
	return s
}

func (t *Trie) PrefixSearchWithElem(pre string) []strKV {
	node := findNode(t.root, []rune(pre))
	if node == nil {
		return nil
	}

	return collectWithElem(node)
}

type strKV struct {
	Key   string
	Value interface{}
}

func (t *Trie) Add(key string, data interface{}) {
	top := t.root
	runes := []rune(key)
	runesByLen := len(runes)

	for index, r := range runes {
		if nByQuery, ok := top.children[r]; ok {
			if index == (runesByLen - 1) {
				if !nByQuery.term {
					t.size++
				}
				nByQuery.path = key
				nByQuery.term = true
				nByQuery.meta = data
			}
			top = nByQuery
		} else {
			if index == (runesByLen - 1) {
				t.size++
				top = top.newChild(r, key, data, true)
			} else {
				top = top.newChild(r, "", nil, false)
			}
		}
	}
}

func (t *Trie) Find(key string) *node {
	nByQuery := findNode(t.root, []rune(key))
	if nByQuery != nil && nByQuery.term {
		return nByQuery
	}
	return nil
}

func (t *Trie) Remove(key string) {
	nByQuery := findNode(t.root, []rune(key))
	if nByQuery != nil && nByQuery.term {
		t.size--

		nByQuery.path = ""
		nByQuery.term = false
		nByQuery.meta = nil

		if len(nByQuery.parent.children) > 1 && len(nByQuery.children) == 0 {
			delete(nByQuery.parent.children, nByQuery.r)
			nByQuery.parent = nil
			nByQuery.children = nil
		}
	}
}

func (n *node) newChild(r rune, path string, meta interface{}, term bool) *node {
	nByNew := &node{
		r:        r,
		path:     path,
		term:     term,
		meta:     meta,
		parent:   n,
		children: make(map[rune]*node),
	}
	n.children[nByNew.r] = nByNew
	return nByNew
}
func (n *node) GetStrKV() strKV {
	return strKV{
		Key:   n.path,
		Value: n.meta,
	}
}

func findNode(n *node, runes []rune) *node {
	if len(runes) == 0 {
		return n
	}

	nByQuery, ok := n.children[runes[0]]
	if !ok {
		return nil
	}

	var nrunes []rune
	if len(runes) > 1 {
		nrunes = runes[1:]
	} else {
		nrunes = runes[0:0]
	}

	return findNode(nByQuery, nrunes)
}
