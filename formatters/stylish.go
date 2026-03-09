package formatters

import (
	"fmt"
	"sort"
	"strings"

	"code/models"
)

const indentSize = 4

func indent(depth int) string {
	n := depth*indentSize - 2
	if n < 0 {
		n = 0
	}
	return strings.Repeat(" ", n)
}

func FormatStylish(nodes []*models.Node) string {
	result, _ := render(nodes, 1)
	return result
}

func render(nodes []*models.Node, depth int) (string, error) {
	base := indent(depth)
	closeIndent := strings.Repeat(" ", (depth-1)*indentSize)

	var b strings.Builder
	b.WriteString("{\n")

	for _, node := range nodes {
		switch node.Type {
		case "nested":
			childStr, _ := render(node.Children, depth+1)
			b.WriteString(fmt.Sprintf("%s  %s: %s\n", base, node.Key, childStr))
		case "unchanged":
			b.WriteString(fmt.Sprintf("%s  %s: %s\n", base, node.Key, stringify(node.OldValue, depth+1)))
		case "removed":
			val := stringify(node.OldValue, depth+1)
			if val == "" {
				b.WriteString(fmt.Sprintf("%s- %s:\n", base, node.Key))
			} else {
				b.WriteString(fmt.Sprintf("%s- %s: %s\n", base, node.Key, val))
			}
		case "added":
			val := stringify(node.NewValue, depth+1)
			b.WriteString(fmt.Sprintf("%s+ %s: %s\n", base, node.Key, val))
		case "changed":
			oldVal := stringify(node.OldValue, depth+1)
			newVal := stringify(node.NewValue, depth+1)
			if oldVal == "" {
				b.WriteString(fmt.Sprintf("%s- %s:\n", base, node.Key))
			} else {
				b.WriteString(fmt.Sprintf("%s- %s: %s\n", base, node.Key, oldVal))
			}
			b.WriteString(fmt.Sprintf("%s+ %s: %s\n", base, node.Key, newVal))
		}
	}

	b.WriteString(closeIndent + "}")
	return b.String(), nil
}

func stringify(value interface{}, depth int) string {
	if value == nil {
		return "null"
	}

	if m, ok := value.(map[string]interface{}); ok {
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var b strings.Builder
		b.WriteString("{\n")
		valueIndent := strings.Repeat(" ", depth*indentSize)
		for _, k := range keys {
			b.WriteString(fmt.Sprintf("%s%s: %s\n", valueIndent, k, stringify(m[k], depth+1)))
		}
		b.WriteString(strings.Repeat(" ", (depth-1)*indentSize) + "}")
		return b.String()
	}

	switch v := value.(type) {
	case string:
		if v == "" {
			return ""
		}
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}
