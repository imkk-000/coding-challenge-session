package main

import "fmt"

func CounterName(names []string) []string {
	currentMap := make(map[string]int)
	for index, name := range names {
		if _, existing := currentMap[name]; existing {
			names[index] = fmt.Sprintf("%s%d", name, currentMap[name])
			currentMap[name]++
		} else {
			currentMap[name] = 1
		}
	}
	return names
}
