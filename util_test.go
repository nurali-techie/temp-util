package util

import "testing"

func TestUserName(t *testing.T) {
	want := "nurali-virani"
	if got := UserName(" Nurali ", " Virani "); got != want {
		t.Errorf("UserName() = %q, want = %q", got, want)
	}
}

func TestGeneratePassword(t *testing.T) {
	if got := GeneratePassword(); len(got) == 0 {
		t.Errorf("Invalid password generated")
	}
}
