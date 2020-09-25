package sorted_set

import "testing"

func TestStringBuilder(t *testing.T) {
	cases := map[string]struct {
		input    []string
		expected []string
	}{
		"empty": {
			input:    []string{},
			expected: []string{},
		},
		"1 value": {
			input:    []string{"c"},
			expected: []string{"c"},
		},
		"3 values in order": {
			input:    []string{"c", "d", "e"},
			expected: []string{"c", "d", "e"},
		},
		"3 values out of order": {
			input:    []string{"e", "d", "b"},
			expected: []string{"b", "d", "e"},
		},
		"add duplicate values": {
			input:    []string{"e", "e", "e", "a"},
			expected: []string{"a", "e"},
		},
	}
	for caseName, c := range cases {
		t.Run(caseName, func(t *testing.T) {
			actual := NewString(c.input...).Sort()
			if len(actual) != len(c.expected) {
				t.Errorf("expected %d items, but got %d", len(c.expected), len(actual))
				return
			}
			for i, s := range c.expected {
				if s != actual[i] {
					t.Errorf("expected item %d to be '%s', but got '%s'", i, s, actual[i])
				}
			}
		})
	}
}
