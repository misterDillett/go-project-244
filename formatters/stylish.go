package formatters

import (
	"fmt"
	"sort"
	"strings"

	"code/models"
)

func FormatStylish(nodes []*models.Node, depth ...int) string {
	currentDepth := 0
	if len(depth) > 0 {
		currentDepth = depth[0]
	}

	indent := strings.Repeat("    ", currentDepth)
	result := make([]string, 0, len(nodes))

	for _, node := range nodes {
		switch node.Type {
		case "added":
			result = append(result, fmt.Sprintf("%s  + %s: %s", indent, node.Key, stringify(node.NewValue, currentDepth+1)))
		case "removed":
			val := stringify(node.OldValue, currentDepth+1)
			if val == "" {
				result = append(result, fmt.Sprintf("%s  - %s:", indent, node.Key))
			} else {
				result = append(result, fmt.Sprintf("%s  - %s: %s", indent, node.Key, val))
			}
		case "unchanged":
			result = append(result, fmt.Sprintf("%s    %s: %s", indent, node.Key, stringify(node.OldValue, currentDepth+1)))
		case "changed":
			oldVal := stringify(node.OldValue, currentDepth+1)
			newVal := stringify(node.NewValue, currentDepth+1)
			if oldVal == "" {
				result = append(result, fmt.Sprintf("%s  - %s:", indent, node.Key))
			} else {
				result = append(result, fmt.Sprintf("%s  - %s: %s", indent, node.Key, oldVal))
			}
			result = append(result, fmt.Sprintf("%s  + %s: %s", indent, node.Key, newVal))
		case "nested":
			nested := FormatStylish(node.Children, currentDepth+1)
			result = append(result, fmt.Sprintf("%s    %s: %s", indent, node.Key, nested))
		}
	}

	if currentDepth == 0 {
		return "{\n" + strings.Join(result, "\n") + "\n}"
	}
	return "{\n" + strings.Join(result, "\n") + "\n" + indent + "}"
}

func stringify(value interface{}, depth int) string {
	switch v := value.(type) {
	case map[string]interface{}:
		keys := make([]string, 0, len(v))
		for key := range v {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		nodes := make([]*models.Node, 0, len(v))
		for _, key := range keys {
			nodes = append(nodes, &models.Node{
				Key:      key,
				Type:     "unchanged",
				OldValue: v[key],
			})
		}
		return FormatStylish(nodes, depth)
	case string:
		if v == "" {
			return ""
		}
		return v
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%v", v)
	}
}
