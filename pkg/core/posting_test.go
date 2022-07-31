package core

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/numary/ledger/pkg/core/monetary"
)

func TestReverseMultiple(t *testing.T) {
	p := Postings{
		{
			Source:      "world",
			Destination: "users:monetary.NewInt(00)1",
			Amount:      monetary.NewInt(100),
			Asset:       "COIN",
		},
		{
			Source:      "users:monetary.NewInt(00)1",
			Destination: "payments:monetary.NewInt(00)1",
			Amount:      monetary.NewInt(100),
			Asset:       "COIN",
		},
	}

	expected := Postings{
		{
			Source:      "payments:monetary.NewInt(00)1",
			Destination: "users:monetary.NewInt(00)1",
			Amount:      monetary.NewInt(100),
			Asset:       "COIN",
		},
		{
			Source:      "users:monetary.NewInt(00)1",
			Destination: "world",
			Amount:      monetary.NewInt(100),
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
			Destination: "users:monetary.NewInt(00)1",
			Amount:      monetary.NewInt(100),
			Asset:       "COIN",
		},
	}

	expected := Postings{
		{
			Source:      "users:monetary.NewInt(00)1",
			Destination: "world",
			Amount:      monetary.NewInt(100),
			Asset:       "COIN",
		},
	}

	p.Reverse()

	if diff := cmp.Diff(expected, p); diff != "" {
		t.Errorf("Reverse() mismatch (-want +got):\n%s", diff)
	}
}
