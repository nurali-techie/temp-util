package util

import "testing"

func TestUserName(t *testing.T) {
	want := "nurali-virani-b"
	if got := UserName(" Nurali ", " Virani ", "Barkatali"); got != want {
		t.Errorf("UserName() = %q, want = %q", got, want)
	}
}

func TestGeneratePassword(t *testing.T) {
	if got := GeneratePassword(); len(got) == 0 {
		t.Errorf("Invalid password generated")
	}
}

func TestGenerateEmail(t *testing.T) {
	want := "nurali.virani.b@mycompany.com"
	if got := GenerateEmail("nurali-virani-b"); got != want {
		t.Errorf("GenerateEmail() = %q, want = %q", got, want)
	}
}
