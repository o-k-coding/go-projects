package main

import "testing"

func TestTrie (t *testing.T) {
	t.Run("TestInsertAndSearchSimple", func(t *testing.T) {
		testTrie := InitTrie()

		testTrie.Insert("aragorn")
		result := testTrie.Search("aragorn")

		if result == false {
			t.Error("Expected to find aragorn after inserting aragorn")
		}

		result = testTrie.Search("arathorn")

		if result == true {
			t.Error("Expected not to find arathorn after inserting aragorn")
		}
	})

	t.Run("TestInsertAndSearchMany", func(t *testing.T) {
		toAdd := []string{
			"aragorn",
			"aragon",
			"argon",
			"eragon",
			"oregon",
			"oregano",
			"oreo",
		}
		testTrie := InitTrie()
		for _, word := range toAdd {
			testTrie.Insert(word)
			result := testTrie.Search(word)
			if result == false {
				t.Errorf("Expected to find %s after inserting", word)
			}
		}

		toSearch := []string{
			"algolia",
			"arathon",
			"arg",
			"ergon",
			"ore",
			"oreono",
		}

		for _, word := range toSearch {
			result := testTrie.Search(word)
			if result == true {
				t.Errorf("Expected not to find %s", word)
			}
		}
	})
}
