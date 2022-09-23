package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")

		assertError(t, err, nil)
		want := "this is just a test"
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		assertStrings(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "new test")

		assertError(t, err, ErrWordExists)
		want := "this is just a test"
		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("should find added word:", err)
		}
		assertStrings(t, got, want)
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	err := dictionary.Update("test", "new definition")
	assertError(t, err, nil)

	want := "new definition"
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, want)
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	dictionary.Delete("test")

	_, err := dictionary.Search("test")
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", "test")
	}
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to get an error.")
	}
}
