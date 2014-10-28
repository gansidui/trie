package trie

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestInsert(t *testing.T) {
	tr := New()
	if tr.Size() != 0 {
		t.Fatal()
	}

	tr.Insert("a", "a")
	tr.Insert("a", "a")
	tr.Insert("ab", "ab")
	tr.Insert("cd", "cd")
	tr.Insert("abcd", "abcd")
	tr.Insert("你好", "你好")
	tr.Insert("世界", "世界")
	tr.Insert("helloworld", "helloworld")

	if tr.Size() != 7 {
		t.Fatal()
	}

	flag, value, index := tr.Find("ab")
	if flag != true || value.(string) != "ab" || index != 2 {
		t.Fatal()
	}
	flag, value, index = tr.FindByRunes([]rune{'a', 'b'})
	if flag != true || value.(string) != "ab" || index != 2 {
		t.Fatal()
	}

	flag, value, index = tr.Find("cde")
	if flag != false || value != nil || index != 2 {
		t.Fatal()
	}
	flag, value, index = tr.FindByRunes([]rune{'c', 'd', 'e'})
	if flag != false || value != nil || index != 2 {
		t.Fatal()
	}

	flag, value, index = tr.Find("abcde")
	if flag != false || value != nil || index != 4 {
		t.Fatal()
	}
	flag, value, index = tr.FindByRunes([]rune{'a', 'b', 'c', 'd', 'e'})
	if flag != false || value != nil || index != 4 {
		t.Fatal()
	}

	flag, value, index = tr.Find("你好")
	if flag != true || value.(string) != "你好" || index != 6 {
		t.Fatal()
	}
	flag, value, index = tr.FindByRunes([]rune{'你', '好'})
	if flag != true || value.(string) != "你好" || index != 2 {
		t.Fatal()
	}

	flag, value, index = tr.Find("世界你好")
	if flag != false || value != nil || index != 6 {
		t.Fatal()
	}
	flag, value, index = tr.FindByRunes([]rune{'世', '界', '你', '好'})
	if flag != false || value != nil || index != 2 {
		t.Fatal()
	}

	flag, value, index = tr.Find("hello")
	if flag != false || value != nil || index != 0 {
		t.Fatal()
	}
	flag, value, index = tr.FindByRunes([]rune{'h', 'e', 'l', 'l', 'o'})
	if flag != false || value != nil || index != 0 {
		t.Fatal()
	}
}

func TestDelete(t *testing.T) {
	tr := New()
	tr.Insert("a", "a")
	tr.Insert("a", "a")
	tr.Insert("ab", "ab")
	tr.Insert("abcd", "abcd")
	tr.Insert("b", "b")
	tr.Insert("bc", "bc")

	if tr.Size() != 5 {
		t.Fatal()
	}

	flag, value, index := tr.Find("abc")
	if flag != false || value != nil || index != 2 {
		t.Fatal()
	}

	tr.Delete("ab")
	if tr.Size() != 4 {
		t.Fatal()
	}

	flag, value, index = tr.Find("abc")
	if flag != false || value != nil || index != 1 {
		t.Fatal()
	}

	tr.Delete("a")
	if tr.Size() != 3 {
		t.Fatal()
	}
}

func TestPrefixMatch(t *testing.T) {
	tr := New()
	tr.Insert("h", "h")
	tr.Insert("hi", "hi")
	tr.Insert("hello", "hello")
	tr.Insert("hey", "hey")
	tr.Insert("hei", "hei")
	tr.Insert("helloworld", "helloworld")
	tr.Insert("我", "我")
	tr.Insert("我是", "我是")
	tr.Insert("我是gan", "我是gan")
	tr.Insert("我是gansidui", "我是gansidui")

	if tr.Size() != 10 {
		t.Fatal()
	}

	ret := tr.PrefixMatch("")
	if len(ret) != tr.Size() {
		t.Fatal()
	}

	ret = tr.PrefixMatch("a")
	if len(ret) != 0 {
		t.Fatal()
	}

	ret = tr.PrefixMatch("h")
	if len(ret) != 6 {
		t.Fatal()
	}

	ret = tr.PrefixMatch("hello")
	if len(ret) != 2 || ret[0].(string) != "hello" || ret[1].(string) != "helloworld" {
		t.Fatal()
	}

	if len(tr.PrefixMatch("我")) != 4 || len(tr.PrefixMatch("我是")) != 3 {
		t.Fatal()
	}

	ret = tr.PrefixMatch("我是gan")
	if len(ret) != 2 || ret[0].(string) != "我是gan" || ret[1].(string) != "我是gansidui" {
		t.Fatal()
	}
	ret = tr.PrefixMatchByRunes([]rune{'我', '是', 'g', 'a', 'n'})
	if len(ret) != 2 || ret[0].(string) != "我是gan" || ret[1].(string) != "我是gansidui" {
		t.Fatal()
	}

	ret = tr.PrefixMatch("哈哈")
	if len(ret) != 0 {
		t.Fatal()
	}

	tr.Delete("h")
	if len(tr.PrefixMatch("h")) != 5 {
		t.Fatal()
	}

	ret = tr.PrefixMatch("")
	if len(ret) != 9 {
		t.Fatal()
	}
}

func BenchmarkFind(b *testing.B) {
	br := New()
	for i := 0; i < b.N; i++ {
		br.Insert("key"+strconv.Itoa(i), "value")
	}

	for i := 0; i < b.N; i++ {
		br.Find("key" + strconv.Itoa(i))
	}
}

func BenchmarkDelete(b *testing.B) {
	br := New()
	for i := 0; i < b.N; i++ {
		br.Insert("key"+strconv.Itoa(i), "value")
	}

	for i := 0; i < b.N; i++ {
		br.Delete("key" + strconv.Itoa(i))
	}
}

func BenchmarkPrefixMatch(b *testing.B) {
	br := New()
	for i := 0; i < b.N; i++ {
		a := rand.Intn(10)
		b := rand.Intn(10)
		c := rand.Intn(10)
		prefix := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c)
		br.Insert(prefix+"key"+strconv.Itoa(i), "value")
	}

	for i := 0; i < b.N; i++ {
		a := rand.Intn(10)
		b := rand.Intn(10)
		c := rand.Intn(10)
		prefix := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c)
		br.PrefixMatch(prefix + "key" + strconv.Itoa(i))
	}
}
