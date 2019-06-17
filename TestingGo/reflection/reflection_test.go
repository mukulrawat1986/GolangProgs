package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	// make table test

	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			Name: "Struct with one string field",
			Input: struct{
				Name string
			}{
				Name: "Chris",
			},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with two string fields",
			Input: struct {
				Name string
				City string
			}{
				Name: "Chris",
				City: "London",
			},
			ExpectedCalls: []string{
				"Chris",
				"London",
			},
		},
		{
			Name: "Struct with non-string field",
			Input: struct{
				Name string
				Age int
			}{
				Name: "Chris",
				Age: 33,
			},
			ExpectedCalls: []string{
				"Chris",
			},
		},
		{
			Name: "Nested fields",
			Input: struct{
				Name string
				Profile struct {
					Age int
					City string
				}
			}{
				Name: "Chris",
				Profile: struct {
					Age int
					City string
				}{
					Age: 33,
					City: "London",
				},
			},
			ExpectedCalls: []string{
				"Chris",
				"London",
			},
		},
		{
			Name: "Pointers to things",
			Input: &struct {
				Name string
				Profile struct {
					Age int
					City string
				}
			}{
				Name: "Chris",
				Profile: struct {
					Age int
					City string
				}{
					Age: 33,
					City: "London",
				},
			},
			ExpectedCalls: []string{
				"Chris",
				"London",
			},
		},
		{
			Name: "slices",
			Input: []struct{
					Age int
					City string
			}{
				{
					Age: 33,
					City: "London",
				},
				{
					Age: 34,
					City: "Reykjavik",
				},
			},
			ExpectedCalls: []string{
				"London",
				"Reykjavik",
			},
		},
		{
			Name: "Arrays",
			Input: [2]struct {
				Age int
				City string
			}{
				{
					Age: 33,
					City: "London",
				},
				{
					Age: 34,
					City: "Reykjavik",
				},
			},
			ExpectedCalls: []string{
				"London",
				"Reykjavik",
			},
		},
	}


	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string){
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got '%v' want '%v'", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string){
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected '%v' to contain '%s' but it didn't", haystack, needle)
	}
}
