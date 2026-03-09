package formatters

import (
	"code/models"
	"encoding/json"
)

func FormatJSON(diff []*models.Node) (string, error) {
	tree := buildTree(diff)
	result, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func buildTree(nodes []*models.Node) map[string]interface{} {
	result := make(map[string]interface{})

	for _, node := range nodes {
		switch node.Type {
		case "nested":
			result[node.Key] = buildTree(node.Children)

		case "changed":
			result[node.Key] = map[string]interface{}{
				"oldValue": node.OldValue,
				"newValue": node.NewValue,
			}

		case "added":
			result[node.Key] = node.NewValue

		case "removed":
			result[node.Key] = node.OldValue

		case "unchanged":
			result[node.Key] = node.OldValue
		}
	}
	return result
}
