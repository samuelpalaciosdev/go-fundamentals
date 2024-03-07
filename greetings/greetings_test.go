package greetings

import "testing"

func TestHelloName(t *testing.T) {
	want := "Hi, test. Welcome!"
	got := Hello([]string{"test"})
	if want != got {
		t.Errorf("Wanted %s, got %s", want, got)
	}
}
