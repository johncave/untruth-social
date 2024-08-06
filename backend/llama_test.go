package main

import (
	"testing"
)

func TestModeration(t *testing.T) {
	allowed, err := ModeratePost("I love llamas")
	if err != nil {
		t.Fatal(err)
	}
	if !allowed {
		t.Fatal("'I love llamas' should be allowed")
	}

	allowed, err = ModeratePost("I hate llamas")
	if err != nil {
		t.Fatal(err)
	}
	if !allowed {
		t.Fatal("'I hate llamas' should be allowed")
	}

	allowed, err = ModeratePost("Fuck all of yall")
	if err != nil {
		t.Fatal(err)
	}
	if allowed {
		t.Fatal("'Fuck all of yall' should not be allowed")
	}

}
