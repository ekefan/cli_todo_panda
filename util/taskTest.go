package util

import (
	
	"math/rand"
)

func RandomDesc() string {
	descriptions := []string{
		"This is a test decripton",
		"Another task up",
		"GSD",
		"Cook",
		"Clean up",
	}
	descIdx := rand.Intn(len(descriptions))
	return descriptions[descIdx]
}

func RandomPriority() string {
	priorities := []string{
		"high",
		"low",
		"none",
	}
	priorityIdx := rand.Intn(len(priorities))
	return priorities[priorityIdx]
}