package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	t1 := New()

	t1.Add("foobar", struct{}{})
	t1.Remove("fooba")
	t1.Add("foobar", struct{}{})
	if t1.Size() != 1 {
		t.Fatal()

	}

	t1.Add("foo", struct{}{})
	t1.Add("bar", struct{}{})
	t1.Add("barb", struct{}{})
	t1.Add("barfoo", struct{}{})
	if t1.Find("fo") != nil {
		t.Fatal()

	}
	if t1.Find("foo") == nil {
		t.Fatal()
	}
	t1.Add("ba", struct{}{})
	if t1.Size() != 6 {
		t.Fatal()
	}
	t1.Remove("b")
	t1.Remove("ba")
	if t1.Size() != 5 {
		t.Fatal()
	}

	{
		// foobar
		// foo
		// bar
		// barb
		// barfoo
		if len(t1.PrefixSearch("f")) != 2 {
			t.Fatal()
		}
		if len(t1.PrefixSearch("")) != 5 {
			t.Fatal()
		}
		if len(t1.PrefixSearch("bar")) != 3 {
			t.Fatal()
		}

		if len(t1.FuzzySearch("f")) != 3 {
			t.Fatal()
		}
		if len(t1.FuzzySearch("br")) != 4 {
			t.Fatal()
		}
		if len(t1.FuzzySearch("fo")) != 3 {
			t.Fatal()
		}
	}

}
