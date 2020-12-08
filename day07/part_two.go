package daySeven

import (
	"regexp"
	"strconv"
	"strings"
)

type BagPair struct {
	innerBag string
	outerBag string
}

func generateWeightedGraph(rules []string) (map[string][]string, map[BagPair]int) {
	graph := make(map[string][]string)
	weightedGraph := make(map[BagPair]int)
	for _, rule := range rules {
		rawStringAfterSplit := strings.Split(rule, " bags contain ")

		outerBag := rawStringAfterSplit[0]
		pattern := regexp.MustCompile(" (bags?)((, )|(\\.))")
		innerBags := pattern.Split(rawStringAfterSplit[1], -1)
		if _, ok := graph[outerBag]; !ok {
			graph[outerBag] = make([]string, 0)

		}
		for _, innerBagWithAmt := range innerBags {
			if innerBagWithAmt == "" || innerBagWithAmt == "no other" {
				continue
			}
			innerBag := innerBagWithAmt[2:]
			amount, _ := strconv.Atoi(string(innerBagWithAmt[0]))
			weightedGraph[BagPair{outerBag, innerBag}] = amount
			if _, ok := graph[outerBag]; !ok {
				graph[outerBag] = make([]string, 0)
			}

			graph[outerBag] = append(graph[outerBag], innerBag)
		}
	}
	return graph, weightedGraph
}

func PartTwo(rules []string) int {
	graph, weightedGraph := generateWeightedGraph(rules)

	frontier := append(make([]string, 0), "shiny gold")
	frontierAugmentation := append(make([]int, 0), 1)
	found := make(map[string]bool)

	var u string
	var t, s int
	count := 0
	for len(frontier) > 0 {
		// pop from front of frontier
		u, frontier = frontier[0], frontier[1:]
		t, frontierAugmentation = frontierAugmentation[0], frontierAugmentation[1:]

		for _, v := range graph[u] {
			if _, ok := found[v]; !ok {
				found[v] = true
			}
			s = t * weightedGraph[BagPair{u, v}]
			count += s
			frontier = append(frontier, v)
			frontierAugmentation = append(frontierAugmentation, s)
		}
	}
	return count
}
