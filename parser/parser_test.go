package parser

import (
	"fmt"
	"testing"
)

func errorMsg(t *testing.T, expected string, actual string) {
	t.Errorf("expected %s, got %s", expected, actual)
}

func TestParser_Parse(t *testing.T) {
	type testCase struct {
		src    string
		expect string
	}

	cases := []testCase{
		{"-1", "-1"},
		{"+1", "+1"},
		{"1 + 2", "(1 + 2)"},
		{"1 + 2 * 3", "(1 + (2 * 3))"},
		{"1 + 2 * 3 - 1", "((1 + (2 * 3)) - 1)"},
		{"1 / 2 * 3 - 1", "(((1 / 2) * 3) - 1)"},
		{"3 * 5 % 2 + 1", "((3 * (5 % 2)) + 1)"},
		{"( 1 + 2 ) * 3", "((1 + 2) * 3)"},
		{"5 ^ 4 ^ 9", "(5 ^ (4 ^ 9))"},
	}

	for _, tcase := range cases {
		p := NewParser(tcase.src)
		actual := fmt.Sprint(p.Parse())
		if actual != tcase.expect {
			errorMsg(t, tcase.expect, actual)
		}
	}
}
