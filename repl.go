package main

import (
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(strings.TrimSpace(text))
	return strings.Fields(text)
}
