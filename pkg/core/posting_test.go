package core

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReverseMultiple(t *testing.T) {
	p := Postings{
		{
			Source:      "world",
			Destination: "users:NewMonetaryInt(00)1",
			Amount:      NewMonetaryInt(100),
			Asset:       "COIN",
		},
		{
			Source:      "users:NewMonetaryInt(00)1",
			Destination: "payments:NewMonetaryInt(00)1",
			Amount:      NewMonetaryInt(100),
			Asset:       "COIN",
		},
	}

	expected := Postings{
		{
			Source:      "payments:NewMonetaryInt(00)1",
			Destination: "users:NewMonetaryInt(00)1",
			Amount:      NewMonetaryInt(100),
			Asset:       "COIN",
		},
		{
			Source:      "users:NewMonetaryInt(00)1",
			Destination: "world",
			Amount:      NewMonetaryInt(100),
			Asset:       "COIN",
		},
	}

	p.Reverse()

	if diff := cmp.Diff(expected, p); diff != "" {
		t.Errorf("Reverse() mismatch (-want +got):\n%s", diff)
	}
}

func TestReverseSingle(t *testing.T) {
	p := Postings{
		{
			Source:      "world",
			Destination: "users:NewMonetaryInt(00)1",
			Amount:      NewMonetaryInt(100),
			Asset:       "COIN",
		},
	}

	expected := Postings{
		{
			Source:      "users:NewMonetaryInt(00)1",
			Destination: "world",
			Amount:      NewMonetaryInt(100),
			Asset:       "COIN",
		},
	}

	p.Reverse()

	if diff := cmp.Diff(expected, p); diff != "" {
		t.Errorf("Reverse() mismatch (-want +got):\n%s", diff)
	}
}
