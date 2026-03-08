package formatters

import (
	"fmt"
	"strings"

	"code/models"
)

func FormatPlain(nodes []*models.Node, path string) string {
	result := make([]string, 0, len(nodes))

	for _, node := range nodes {
		fullPath := node.Key
		if path != "" {
			fullPath = path + "." + node.Key
		}

		switch node.Type {
		case "added":
			result = append(result, fmt.Sprintf("Property '%s' was added with value: %s", fullPath, formatValue(node.NewValue)))
		case "removed":
			result = append(result, fmt.Sprintf("Property '%s' was removed", fullPath))
		case "changed":
			result = append(result, fmt.Sprintf("Property '%s' was updated. From %s to %s", fullPath, formatValue(node.OldValue), formatValue(node.NewValue)))
		case "nested":
			nested := FormatPlain(node.Children, fullPath)
			if nested != "" {
				result = append(result, nested)
			}
		}
	}

	return strings.Join(result, "\n")
}
