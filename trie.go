package trie

type trieNode struct {
	flag  bool               // 该节点是不是一个词的终点
	value interface{}        // 存储任意值
	child map[rune]*trieNode // 孩子节点
}

func newTrieNode() *trieNode {
	return &trieNode{
		flag:  false,
		value: nil,
		child: make(map[rune]*trieNode),
	}
}

type Trie struct {
	root *trieNode
	size int
}

func New() *Trie {
	return &Trie{
		root: newTrieNode(),
		size: 0,
	}
}

// 得到Trie上的词的总数
func (this *Trie) Size() int {
	return this.size
}

// 插入一个词(key, value)，会覆盖相同的key
func (this *Trie) Insert(key string, value interface{}) {
	curNode := this.root
	for _, v := range key {
		if curNode.child[v] == nil {
			curNode.child[v] = newTrieNode()
		}
		curNode = curNode.child[v]
	}

	if !curNode.flag {
		this.size++
		curNode.flag = true
	}
	curNode.value = value
}

// 删除一个词(key)，删除成功返回true
func (this *Trie) Delete(key string) bool {
	curNode := this.root
	preNode := this.root
	var ru rune
	for _, v := range key {
		if curNode.child[v] == nil {
			return false
		}
		preNode = curNode
		curNode = curNode.child[v]
		ru = v
	}

	// 若是叶子节点，则真正删除，否则懒惰删除
	if len(curNode.child) == 0 {
		delete(preNode.child, ru)
	} else {
		curNode.flag = false
		curNode.value = nil
	}

	this.size--

	return true
}

// 查找一个词(key)
// flag为true时表示存在该词(key)， 否则表示不存在该词
// value为该词(key)所在节点的value值
// index表示该词(key)所在路径上最长的一个词的最后一个rune的末尾位置，
// 比如：词典中有一个词："你好"，key为"你好"，那么index为6
// 比如：词典中有一个词："世界"，key为"世界你好"，那么index为6
// 比如：词典中有一个词："hello"，key为"helloworld"，那么index为5
// 比如：词典中有一个词："helloworld"，key为"hello"，那么index为0
func (this *Trie) Find(key string) (flag bool, value interface{}, index int) {
	node, i := this.findNode(key)
	if node == nil {
		return false, nil, i
	} else {
		return node.flag, node.value, i
	}
}

// 功能跟Find完全一样，只是参数在外面就已经将string拆分成了[]rune
// 返回的index为[]rune的下标：
// 比如：词典中有一个词："世界"，key为"世界你好"，那么index为2
func (this *Trie) FindByRunes(runes []rune) (flag bool, value interface{}, index int) {
	node, i := this.findNodeByRunes(runes)
	if node == nil {
		return false, nil, i
	} else {
		return node.flag, node.value, i
	}
}

// 匹配出所有前缀为key的词所在节点的value值
func (this *Trie) PrefixMatch(key string) []interface{} {
	node, _ := this.findNode(key)
	if node != nil {
		return this.walk(node)
	}
	return []interface{}{}
}

// 功能跟PrefixMatch完全一样，只是参数在外面就已经将string拆分成了[]rune
func (this *Trie) PrefixMatchByRunes(runes []rune) []interface{} {
	node, _ := this.findNodeByRunes(runes)
	if node != nil {
		return this.walk(node)
	}
	return []interface{}{}
}

// 遍历
func (this *Trie) walk(node *trieNode) (ret []interface{}) {
	if node.flag {
		ret = append(ret, node.value)
	}
	for _, v := range node.child {
		ret = append(ret, this.walk(v)...)
	}
	return
}

// 查找一个词(key)所在的节点
// node不为空且node.flag为true时表示存在该词(key)， 否则表示不存在该词
// index表示该词(key)所在路径上最长的一个词的最后一个rune的末尾位置，
// 比如：词典中有一个词："你好"，key为"你好"，那么index为6
// 比如：词典中有一个词："世界"，key为"世界你好"，那么index为6
// 比如：词典中有一个词："hello"，key为"helloworld"，那么index为5
// 比如：词典中有一个词："helloworld"，key为"hello"，那么index为0
func (this *Trie) findNode(key string) (node *trieNode, index int) {
	curNode := this.root
	ff := false
	for k, v := range key {
		if ff {
			index = k
			ff = false
		}
		if curNode.child[v] == nil {
			return nil, index
		}
		curNode = curNode.child[v]
		if curNode.flag {
			ff = true
		}
	}

	if curNode.flag {
		index = len(key)
	}

	return curNode, index
}

// 功能跟findNode完全一样，只是参数在外面就已经将string拆分成了[]rune
// 返回的index为[]rune的下标：
// 比如：词典中有一个词："世界"，key为"世界你好"，那么index为2
func (this *Trie) findNodeByRunes(runes []rune) (node *trieNode, index int) {
	curNode := this.root
	ff := false
	for k, v := range runes {
		if ff {
			index = k
			ff = false
		}
		if curNode.child[v] == nil {
			return nil, index
		}
		curNode = curNode.child[v]
		if curNode.flag {
			ff = true
		}
	}

	if curNode.flag {
		index = len(runes)
	}

	return curNode, index
}
