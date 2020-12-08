package daySeven

import (
	"regexp"
	"strings"
)

func extractAllowedBags(rawAllowedBag string) []string {
	var allowedBags []string
	pattern := regexp.MustCompile(" (bags?)((, )|(\\.))")
	otherBags := pattern.Split(rawAllowedBag, -1)
	for _, rawOtherBag := range otherBags {
		if rawOtherBag == "" || rawOtherBag == "no other " {
			continue
		}
		allowedBags = append(allowedBags, rawOtherBag[2:])
	}
	return allowedBags

}

func generateGraph(rules []string) map[string][]string {
	bagsThatMayContainIt := make(map[string][]string)
	for _, rule := range rules {
		rawStringAfterSplit := strings.Split(rule, " bags contain ")

		mainBag := rawStringAfterSplit[0]
		allowedBags := extractAllowedBags(rawStringAfterSplit[1])
		if _, ok := bagsThatMayContainIt[mainBag]; !ok {
			bagsThatMayContainIt[mainBag] = make([]string, 0)
		}
		for _, bag := range allowedBags {
			if _, ok := bagsThatMayContainIt[bag]; !ok {
				bagsThatMayContainIt[bag] = make([]string, 0)
			}
			bagsThatMayContainIt[bag] = append(bagsThatMayContainIt[bag], mainBag)
		}
	}
	return bagsThatMayContainIt
}

func breathFirstSearch(graph map[string][]string) map[string]bool {
	frontier := append(make([]string, 0), "shiny gold")
	found := make(map[string]bool)
	var u string
	for len(frontier) > 0 {
		u, frontier = frontier[0], frontier[1:]

		for _, v := range graph[u] {
			if _, ok := found[v]; !ok {
				found[v] = true
			}
			frontier = append(frontier, v)
		}
	}
	return found
}

func PartOne(rules []string) int {
	graph := generateGraph(rules)
	found := breathFirstSearch(graph)
	return len(found)
}
