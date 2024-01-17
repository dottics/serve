package servemux

import (
	"context"
)

func Params(c context.Context) []Param {
	val := c.Value("params")
	if val == nil {
		return []Param{}
	}
	return val.([]Param)
}

func GetParam(c context.Context, name string) string {
	val := c.Value("params")
	if val == nil {
		return ""
	}
	switch v := val.(type) {
	case []Param:
		for _, p := range v {
			if p.Name == name {
				return p.Value
			}
		}
	}
	return ""
}
