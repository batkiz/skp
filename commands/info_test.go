package commands

import "testing"

func TestExtractName(t *testing.T) {
	var extractNameTests = []struct {
		in       string // input
		expected string // expected
	}{
		{
			`C:\soft\Scoop\buckets\backit\bucket\delta.json`,
			"delta",
		},
		{
			`https://github.com/chawyehsu/dorado/blob/master/bucket/aegisub.json`,
			"aegisub",
		},
		{
			`.\skp.json`,
			"skp",
		},
	}

	for _, test := range extractNameTests {
		actual := ExtractName(test.in)
		if actual != test.expected {
			t.Errorf("ExtractName(%v) = %v; expected %v", test.in, actual, test.expected)
		}
	}
}
