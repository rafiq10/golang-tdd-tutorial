package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	t.Run("structures where order matters", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{"Rafa"},
				[]string{"Rafa"},
			},
			{
				"struct with two string fields",
				struct {
					Name string
					City string
				}{"Rafa", "Krzczonow"},
				[]string{"Rafa", "Krzczonow"},
			},
			{
				"struct with non-string field",
				struct {
					Name string
					Age  int
				}{"Rafa", 37},
				[]string{"Rafa"},
			},
			{
				"nested struct scenario",
				Person{"Rafa", Profile{37, "Krzczonow"}},
				[]string{"Rafa", "Krzczonow"},
			},
			{
				"pointer to things",
				&Person{"Rafa", Profile{37, "Krzczonow"}},
				[]string{"Rafa", "Krzczonow"},
			},
			{
				"slices",
				[]Profile{
					{37, "Krzczonow"},
					{25, "Lublin"},
				},
				[]string{"Krzczonow", "Lublin"},
			},
			{
				"arrays",
				[2]Profile{
					{37, "Krzczonow"},
					{25, "Lublin"}},
				[]string{"Krzczonow", "Lublin"},
			},
			{
				"maps",
				map[string]string{"37": "Krzczonow", "25": "Lublin"},
				[]string{"Krzczonow", "Lublin"},
			},
		}

		for _, test := range cases {
			t.Run(test.Name, func(t *testing.T) {
				var got []string

				walk(test.Input, func(input string) {
					got = append(got, input)
				})
				if !reflect.DeepEqual(got, test.ExpectedCalls) {
					t.Errorf("wanted %v but got %v", test.ExpectedCalls, got)
				}
			})
		}

	})

	t.Run("tests for mapr where order does not matter", func(t *testing.T) {
		myMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(myMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("tests for channels", func(t *testing.T) {
		myChan := make(chan Profile)

		go func() {
			myChan <- Profile{32, "Wroclaw"}
			myChan <- Profile{33, "Madrid"}
			close(myChan)
		}()
		var got []string
		want := []string{"Wroclaw", "Madrid"}

		walk(myChan, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanetd: %v but got: %v", want, got)
		}
	})

	t.Run("test for functions", func(t *testing.T) {
		myFunction := func() (Profile, Profile) {
			return Profile{32, "Wroclaw"}, Profile{33, "Madrid"}
		}

		var got []string
		want := []string{"Wroclaw", "Madrid"}

		walk(myFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	})
}

func assertContains(t *testing.T, haystack []string, input string) {
	t.Helper()
	contains := false
	for _, v := range haystack {
		if v == input {
			contains = true
		}
	}

	if !contains {
		t.Errorf("Expected: %+v to contain %q but it didn't", haystack, input)
	}
}
