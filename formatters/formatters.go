package formatters

import (
	"fmt"

	"code/models"
)

func Format(diff []*models.Node, format string) (string, error) {
	switch format {
	case "stylish":
		return FormatStylish(diff), nil
	case "plain":
		return FormatPlain(diff, ""), nil
	case "json":
		return FormatJSON(diff)
	default:
		return "", fmt.Errorf("unknown format: %s", format)
	}
}

func formatValue(value interface{}) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case bool:
		if v {
			return "true"
		}
		return "false"
	case float64:
		return fmt.Sprintf("%v", v)
	case string:
		return fmt.Sprintf("'%s'", v)
	case map[string]interface{}, []interface{}:
		return "[complex value]"
	default:
		return fmt.Sprintf("%v", v)
	}
}
