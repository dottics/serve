package servemux

import (
	"regexp"
	"testing"
)

func TestCompilePathRegex(t *testing.T) {
	tests := []struct {
		name            string
		pattern         string
		expectedPattern string
	}{
		{
			pattern:         "/some/{var1}/",
			expectedPattern: `/some/(?P<var1>[a-zA-Z]+)/`,
		},
		{
			pattern:         "/some/{var1}/{var2}",
			expectedPattern: `/some/(?P<var1>[a-zA-Z]+)/(?P<var2>[a-zA-Z]+)`,
		},
		{
			pattern:         "/some/na{var1}/{var2}",
			expectedPattern: `/some/na(?P<var1>[a-zA-Z]+)/(?P<var2>[a-zA-Z]+)`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pat, _ := CompilePathRegex(test.pattern)
			if pat != test.expectedPattern {
				t.Errorf("expected pattern: %v got %v", test.expectedPattern, pat)
			}
		})
	}
}

func TestParsePath(t *testing.T) {
	tests := []struct {
		name           string
		path           string
		pattern        *regexp.Regexp
		params         []Param
		expectedParams []Param
	}{
		{},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			params := test.params
			xp := ParsePath(test.path, test.pattern, params)
			for i, p := range test.expectedParams {
				if p != xp[i] {
					t.Errorf("expected param: %v got %v", p, params[i])
				}
			}
		})
	}
}
