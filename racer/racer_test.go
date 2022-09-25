package racer

import "testing"

func TestRace(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.co.uk"

	want := fastURL
	got := Racer(slowURL, fastURL)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
