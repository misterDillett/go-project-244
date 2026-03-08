package code

import (
	"sort"

	"code/formatters"
	"code/models"
	"code/parsers"
)

func GenDiff(filepath1, filepath2, format string) (string, error) {
	data1, err := parsers.ParseFile(filepath1)
	if err != nil {
		return "", err
	}

	data2, err := parsers.ParseFile(filepath2)
	if err != nil {
		return "", err
	}

	diff := BuildDiff(data1, data2)

	return formatters.Format(diff, format)
}

func BuildDiff(data1, data2 map[string]interface{}) []*models.Node {
	allKeys := make(map[string]bool)
	for k := range data1 {
		allKeys[k] = true
	}
	for k := range data2 {
		allKeys[k] = true
	}

	keys := make([]string, 0, len(allKeys))
	for k := range allKeys {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var nodes []*models.Node
	for _, key := range keys {
		val1, ok1 := data1[key]
		val2, ok2 := data2[key]

		switch {
		case ok1 && !ok2:
			nodes = append(nodes, &models.Node{
				Key:      key,
				Type:     "removed",
				OldValue: val1,
			})
		case !ok1 && ok2:
			nodes = append(nodes, &models.Node{
				Key:      key,
				Type:     "added",
				NewValue: val2,
			})
		default:
			if isNested(val1, val2) {
				children := BuildDiff(val1.(map[string]interface{}), val2.(map[string]interface{}))
				nodes = append(nodes, &models.Node{
					Key:      key,
					Type:     "nested",
					Children: children,
				})
			} else if val1 != val2 {
				nodes = append(nodes, &models.Node{
					Key:      key,
					Type:     "changed",
					OldValue: val1,
					NewValue: val2,
				})
			} else {
				nodes = append(nodes, &models.Node{
					Key:      key,
					Type:     "unchanged",
					OldValue: val1,
				})
			}
		}
	}

	return nodes
}

func isNested(val1, val2 interface{}) bool {
	_, ok1 := val1.(map[string]interface{})
	_, ok2 := val2.(map[string]interface{})
	return ok1 && ok2
}
