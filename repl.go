package main

import (
	"strings"
)

func cleanInput(text string) []string {

	lowered := strings.ToLower(text)
	splitted := strings.Fields(lowered)

	return splitted
}
