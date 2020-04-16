package mylib

import "testing"

func TestMylib1(t *testing.T) {
    want := "Hello Maru"
    if got := MyLib1(); got != want {
        t.Errorf("Failed = %q %q", got, want)
    }
}
