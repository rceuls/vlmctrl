package main

import "strings"

func convert(input string) []mpcCommand {
	splittedInput := strings.Split(input, ",")
	result := make([]mpcCommand, len(splittedInput))
	for index, cmd := range splittedInput {
		switch cmd {
		case "v_up":
			result[index] = volumeUp
		case "v_down":
			result[index] = volumeDown
		}
	}
	return result
}
