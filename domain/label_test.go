package domain

import "testing"

func TestGetLabelPrefix(t *testing.T) {
	tests := []struct {
		name      string
		label     string
		delimiter string
		expected  string
	}{
		{"No delimiter given", "foo/bar", "", "foo/bar"},
		{"Single colon", "foo:bar", ":", "foo"},
		{"Single slash", "foo/bar", "/", "foo"},
		{"Single underscore", "foo_bar", "_", "foo"},
		{"No delimiter present", "foobar", ".", "foobar"},
		{"Delimiter at start", ":foo", ":", ""},
		{"Multiple colons", "foo:bar:baz", ":", "foo"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getLabelPrefix(tc.label, tc.delimiter)
			if actual != tc.expected {
				t.Errorf("Test '%s': expected %s but got %s", tc.name, tc.expected, actual)
			}
		})
	}
}

func TestRequiredLabel(t *testing.T) {
	tests := []struct {
		name     string
		labels   []string
		rl       RequiredLabel
		expected bool
	}{
		{
			name:     "Label found",
			labels:   []string{"foo/bar", "bar/baz"},
			rl:       RequiredLabel{Prefix: "foo", Separator: "/"},
			expected: true,
		},
		{
			name:     "Label not found",
			labels:   []string{"baz/qux", "quux/corge"},
			rl:       RequiredLabel{Prefix: "foo", Separator: "/"},
			expected: false,
		},
		{
			name:     "Empty label list",
			labels:   []string{},
			rl:       RequiredLabel{Prefix: "foo", Separator: "/"},
			expected: false,
		},
		{
			name:     "Prefix without separator",
			labels:   []string{"foobar", "barbaz"},
			rl:       RequiredLabel{Prefix: "foo", Separator: "/"},
			expected: false,
		},
		{
			name:     "Separator but no matching prefix",
			labels:   []string{"baz/foo/qux", "bar/foo/baz"},
			rl:       RequiredLabel{Prefix: "qux", Separator: "/"},
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.rl.DoExist(tc.labels)
			if actual != tc.expected {
				t.Errorf("%s: Expected %t but got %t", tc.name, tc.expected, actual)
			}
		})
	}
}
