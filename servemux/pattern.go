package servemux

import (
	"fmt"
	"regexp"
	"strings"
)

type Param struct {
	index       int
	patternName string
	name        string
	value       string
}

func CompilePathRegex(pattern string) (string, []Param) {
	re := regexp.MustCompile(`{\w+}`)
	paramNames := re.FindAllString(pattern, -1)
	params := make([]Param, len(paramNames))

	for i, p := range paramNames {
		name := strings.Trim(p, "{}")
		params[i] = Param{
			index:       i,
			name:        name,
			patternName: fmt.Sprintf(`(?P<%s>[a-zA-Z]+)`, name),
		}
	}

	patternRegex := pattern
	for _, p := range params {
		reVal := regexp.MustCompile(fmt.Sprintf("{%s}", p.name))
		patternRegex = reVal.ReplaceAllString(patternRegex, p.patternName)
	}
	// the * indicates that the route is open-ended
	if strings.HasSuffix(pattern, "*") {
		patternRegex = strings.Replace(patternRegex, "*", "", -1)
		patternRegex = fmt.Sprintf(`^%s`, patternRegex)
	} else {
		patternRegex = fmt.Sprintf(`^%s$`, patternRegex)
	}
	return patternRegex, params
}

func ParsePath(path string, rePattern *regexp.Regexp, params []Param) []Param {
	matches := rePattern.FindStringSubmatch(path)
	for i, p := range params {
		subNameIndex := rePattern.SubexpIndex(p.name)
		params[i].value = matches[subNameIndex]
	}
	return params
}
