package main

import (
	"strings"
)

func cleanInput(text string) []string {
    // split by whitespaces
	//lowercase
	//trim leading or trailing whitespace

	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	splitted := strings.Fields(lowered)

	return splitted
}
