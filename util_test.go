package util

import "testing"

func TestUserName(t *testing.T) {
	want := "nurali-virani"
	if got := UserName("Nurali", "Virani"); got != want {
		t.Errorf("UserName() = %q, want = %q", got, want)
	}
}
