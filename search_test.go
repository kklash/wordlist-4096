package wordlist4096

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	type Fixture struct {
		Query      string
		Suffixes   []string
		ExactMatch bool
	}

	fixtures := []*Fixture{
		{
			Query:      "",
			Suffixes:   []string{},
			ExactMatch: false,
		},
		{
			Query:      "aaaaa",
			Suffixes:   []string{},
			ExactMatch: false,
		},
		{
			Query:      "azddee",
			Suffixes:   []string{},
			ExactMatch: false,
		},
		{
			Query:      "briar",
			Suffixes:   []string{""},
			ExactMatch: true,
		},
		{
			Query:      "abandon",
			Suffixes:   []string{""},
			ExactMatch: true,
		},
		{
			Query:      "zoo",
			Suffixes:   []string{"m"},
			ExactMatch: false,
		},
		{
			Query:      "inc",
			Suffixes:   []string{"ense", "h", "ision", "lude", "ome", "ubate"},
			ExactMatch: false,
		},
		{
			Query:      "car",
			Suffixes:   []string{"amel", "bon", "cass", "d", "eful", "go", "ibou", "nival", "ousel", "ry", "sick", "toon", "ve"},
			ExactMatch: false,
		},
		{
			Query:      "emu",
			Suffixes:   []string{"", "late"},
			ExactMatch: true,
		},
		{
			Query:      "ran",
			Suffixes:   []string{"cher", "dom", "sack"},
			ExactMatch: false,
		},
		{
			Query:      "quo",
			Suffixes:   []string{"te"},
			ExactMatch: false,
		},
	}

	for _, fixture := range fixtures {
		result := Search(fixture.Query)
		if !reflect.DeepEqual(result.Suffixes, fixture.Suffixes) {
			t.Errorf(
				"wrong word suffix search results on term %q\nWanted %#v\nGot    %#v",
				fixture.Query, fixture.Suffixes, result.Suffixes,
			)
		}

		if result.ExactMatch != fixture.ExactMatch {
			t.Errorf(
				"expected word search for %q to return ExactMatch=%v, got %v",
				fixture.Query, fixture.ExactMatch, result.ExactMatch,
			)
		}
	}

	for _, word := range WordList {
		result := Search(word)
		if !result.ExactMatch {
			t.Errorf("expected to find exact match for word %q in word list", word)
		}
		if len(result.Suffixes) == 0 {
			t.Errorf("found no matches for word %q", word)
		} else if result.Suffixes[0] != "" {
			t.Errorf("expected first suffix for word %q to be empty string", word)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	queries := []string{
		"car",
		"don",
		"fu",
		"a",
		"income",
		"writ",
		"abandon",
		"zoo",
		"medical",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Search(queries[i%len(queries)])
	}
}
