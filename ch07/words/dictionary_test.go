package words

import "testing"

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}

	given := "test"
	got := d.Search(given)
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
