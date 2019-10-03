package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is just a test",
	}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		assertError(t, err, nil)
		assertString(t, got, "this is just a test")
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
