package formatters

import (
	"fmt"
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
			result = append(result, fmt.Sprintf("%s  - %s: %s", indent, node.Key, stringify(node.OldValue, currentDepth+1)))
		case "unchanged":
			result = append(result, fmt.Sprintf("%s    %s: %s", indent, node.Key, stringify(node.OldValue, currentDepth+1)))
		case "changed":
			result = append(result, fmt.Sprintf("%s  - %s: %s", indent, node.Key, stringify(node.OldValue, currentDepth+1)))
			result = append(result, fmt.Sprintf("%s  + %s: %s", indent, node.Key, stringify(node.NewValue, currentDepth+1)))
		case "nested":
			nested := FormatStylish(node.Children, currentDepth+1)
			result = append(result, fmt.Sprintf("%s    %s: %s", indent, node.Key, nested))
		}
	}

	if currentDepth == 0 {
		return "{\n" + strings.Join(result, "\n") + "\n}"
	}
	return "{\n" + strings.Join(result, "\n") + "\n" + strings.Repeat("    ", currentDepth-1) + "}"
}

func stringify(value interface{}, depth int) string {
	switch v := value.(type) {
	case map[string]interface{}:
		nodes := make([]*models.Node, 0, len(v))
		for key, val := range v {
			nodes = append(nodes, &models.Node{
				Key:      key,
				Type:     "unchanged",
				OldValue: val,
			})
		}
		return FormatStylish(nodes, depth)
	case string:
		return v
	default:
		return formatValue(v)
	}
}
