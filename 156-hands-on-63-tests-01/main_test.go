package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(5, 5) //in terminal run to get PASS status. If incorrect will let you know FAIL
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
