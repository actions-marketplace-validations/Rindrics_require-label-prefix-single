package domain

import "strings"

type RequiredLabel struct {
	prefix    string
	separator string
}

func getLabelPrefix(label string, separator string) string {
	if separator == "" {
		return label
	}

	return strings.Split(label, separator)[0]
}

func (l *RequiredLabel) DoExist(labels []string) bool {
	for _, label := range labels {
		if l.prefix == getLabelPrefix(label, l.separator) {
			return true
		}
	}

	return false
}
