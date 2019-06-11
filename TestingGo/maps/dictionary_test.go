package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{
		"test": "this is just a test",
	}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertString(t, got, want)
}

func assertString(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
