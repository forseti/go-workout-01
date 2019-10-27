package c12reflect

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile Profile
}

type Profile struct {
	Age int
	City string
}

func assertContains(t *testing.T, haystack []string, needle string)  {
	t.Helper()

	contains := false

	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected: %+v to contain %q but it didn't", haystack, needle)
	}
}

func TestWalk(t *testing.T) {
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got[] string

		walk(aMap, func(input string){
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with every other kinds", func(t *testing.T) {
		cases := []struct {
			Name string
			Input interface{}
			ExpectedCalls []string
		}{
			{
				"struct with 1 string field",
				struct {
					Name string
				} { "Chris" },
				[]string{"Chris"},
			},
			{
				"struct with 2 string fields",
				struct {
					Name string
					City string
				}{"Chris", "London"},
				[]string{"Chris", "London"},
			},
			{
				"struct with non string field",
				struct {
					Name string
					Age int
				}{"Chris", 33},
				[]string{"Chris"},
			},
			{
				"nested fields",
				Person {
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"pointers to things",
				&Person {
					"Chris",
					Profile{33, "London"},
				},
				[]string{"Chris", "London"},
			},
			{
				"slices",
				[]Profile{
					{33, "London"},
					{34, "Reykjavik"},
				},
				[]string{"London", "Reykjavik"},
			},
			{
				"Arrays",
				[2]Profile {
					{33, "London"},
					{34, "Reykjavik"},
				},
				[]string{"London", "Reykjavik"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string
				walk(test.Input, func(input string){
					got = append(got, input)
				})

				expected := test.ExpectedCalls
				numOfCalls := len(expected)

				if len(got) != len(expected) {
					t.Errorf("wrong number of function numOfCalls, got %d want %d", len(got), numOfCalls)
				}

				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("got %v want %v", got, expected)
				}
			})
		}
	})
}
