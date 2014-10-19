##trie

~~~ go
package main

import (
	"fmt"
	"github.com/gansidui/trie"
)

func main() {
	tr := trie.New()

	tr.Insert("hello", "world")
	tr.Insert("你好", "value1")
	tr.Insert("你好啊", "value2")
	tr.Insert("你好world", "value3")
	tr.Insert("你好世界", "value4")

	flag, value, index := tr.Find("你好吗")
	fmt.Println(flag, value, index)

	flag, value, index = tr.Find("你好啊")
	fmt.Println(flag, value.(string), index)

	ret := tr.PrefixMatch("你好")
	for i, v := range ret {
		fmt.Println(i, v.(string))
	}

	tr.Delete("hello")

	// output:
	/*
		false <nil> 6
		true value2 9
		0 value1
		1 value2
		2 value3
		3 value4
	*/
}

~~~


##LICENSE

MIT