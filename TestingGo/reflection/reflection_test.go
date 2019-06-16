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
		{
			Name: "Maps",
			Input: map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			ExpectedCalls: []string{
				"Bar",
				"Boz",
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
}
