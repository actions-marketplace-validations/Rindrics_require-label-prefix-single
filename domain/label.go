package domain

import "strings"

type RequiredLabel struct {
	Prefix    string
	Separator string
}

func getLabelPrefix(label string, separator string) string {
	if separator == "" {
		return label
	}

	return strings.Split(label, separator)[0]
}

func (l *RequiredLabel) DoExist(labels Labels) bool {
	for _, label := range labels {
		if l.Prefix == getLabelPrefix(label, l.Separator) {
			return true
		}
	}

	return false
}
